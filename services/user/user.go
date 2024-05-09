package user

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cuixiaojun001/linkhome/common/errdef"
	"github.com/cuixiaojun001/linkhome/common/logger"
	"github.com/cuixiaojun001/linkhome/modules/user/dao"
	"github.com/cuixiaojun001/linkhome/modules/user/model"
	"github.com/cuixiaojun001/linkhome/third_party/sms"
	"github.com/golang-jwt/jwt"
)

const (
	PhoneRegex           = `^1[3456789]\d{9}$`
	JwtSecret            = "NmUzODk2ZGUtYmZjYy0xMWVjLWI5YTctZjQzMGI5YTUwMzQ2aHVp"
	JwtExpiryHours       = 2
	JwtRefreshExpiryDays = 14
)

type IUserService interface {
	// Login 用户登陆
	Login(ctx context.Context, req LoginRequest) (*TokenItem, error)
	// Register 用户注册
	Register(ctx context.Context, req RegisterRequest) (*TokenItem, error)
	// SendSmsCode 发送短信验证码
	SendSmsCode(ctx context.Context, mobile string) error
	// PwdChange 密码修改
	PwdChange(ctx context.Context, userID string, req PwdChangeRequest) (*TokenItem, error)
	// Profile 用户详情信息
	Profile(ctx context.Context, id string) (*UserProfileItem, error)

	// PublishOrUpdateRentalDemand 发布或更新租房需求
	PublishOrUpdateRentalDemand(_ context.Context, id int, req RentalDemandRequest) error
}

type UserService struct{}

var once sync.Once
var houseManager IUserService

func GetUsereManager() IUserService {
	once.Do(func() {
		houseManager = &UserService{}
	})
	return houseManager
}

var _ IUserService = (*UserService)(nil)

func (s *UserService) Login(_ context.Context, req LoginRequest) (*TokenItem, error) {
	// 判断用户账号是手机号还是用户名
	pattern := regexp.MustCompile(PhoneRegex)
	result := pattern.MatchString(req.Account)

	filterParams := map[string]string{"password": req.Password}
	if result { // 手机号
		filterParams["mobile"] = req.Account
	} else { // 用户名
		filterParams["username"] = req.Account
	}

	// 查询用户是否存在
	if user, err := dao.FilterFirst(filterParams); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	} else if err != nil {
		logger.Errorw("GetUserBasicInfo", "err", err)
		return nil, err
	} else {
		return generateUserToken(user, true) // 生成用户token
	}
}

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Refresh  bool   `json:"refresh"`
	jwt.StandardClaims
}

func generateUserToken(user model.UserBasicInfo, withRefreshToken bool) (*TokenItem, error) {
	item := TokenItem{}
	var err error
	// 正常token时效2小时
	now := time.Now()
	expiryTime := now.Add(time.Hour * JwtExpiryHours)
	payload := Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		Refresh:  false,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime.Unix(),
		},
	}
	if item.Token, err = generateJWT(payload); err != nil {
		return &item, err
	}

	// 刷新token时效2周
	if withRefreshToken {
		payload.Refresh = true
		refreshExpiryTime := now.Add(time.Hour * 24 * JwtRefreshExpiryDays)
		payload.ExpiresAt = refreshExpiryTime.Unix()
		if item.RefreshToken, err = generateJWT(payload); err != nil {
			return &item, err
		}
	}

	return &item, nil
}

func generateJWT(payload Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(JwtSecret))
}

func (s *UserService) SendSmsCode(_ context.Context, mobile string) error {
	errCh := make(chan error)
	go sms.Client.SendSmsCode(mobile, errCh)

	err := <-errCh
	if err != nil {
		logger.Errorw("SendSmsCode", "err", err)
		return err
	}
	return nil
}

func (s *UserService) Register(_ context.Context, req RegisterRequest) (*TokenItem, error) {
	if err := verifyUserRegisterInfo(req); err != nil {
		logger.Infow("verifyUserRegisterInfo", "err", err)
		return nil, err
	}
	// 注册用户
	user := model.UserBasicInfo{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
		Mobile:   req.Mobile,
		State:    "normal",
	}

	if err := dao.CreateUserBasic(&user); err != nil {
		logger.Errorw("CreateUserBasic", "err", err)
		return nil, err
	}

	if err := dao.CreateUserProfile(&model.UserProfileInfo{Id: user.ID, Mobile: user.Mobile, State: user.State}); err != nil {
		logger.Errorw("CreateUserProfile", "err", err)
		return nil, err
	}

	// 注册成功保持登陆状态，签发token
	return generateUserToken(user, true)
}

func verifyUserRegisterInfo(item RegisterRequest) error {
	if code := sms.Client.GetSmsCode(item.Mobile); code != item.SmsCode {
		return errdef.SmsCodeErr
	}
	if IsUsernameExist(item.Username) || IsMobileExist(item.Mobile) {
		return errdef.AccountErr
	}

	return nil
}

func IsUsernameExist(username string) bool {
	return dao.IsUsernameExist(username)
}

func IsMobileExist(mobile string) bool {
	return dao.IsMobileExist(mobile)
}

func (s *UserService) PwdChange(_ context.Context, userID string, req PwdChangeRequest) (*TokenItem, error) {
	// 验证旧密码
	user, err := dao.FilterFirst(map[string]string{"id": userID, "password": req.SrcPassword})
	if err != nil {
		return nil, errdef.AccountErr
	}

	if req.NewPassword != req.ConfirmPassword {
		return nil, errdef.CpwdErr
	}

	// 更新密码和token
	if err := dao.UpdateUserPwd(userID, req.NewPassword); err != nil {
		logger.Errorw("UpdateUserPwd", "err", err)
		return nil, err
	}
	// 生成新token
	return generateUserToken(user, true)
}

func (s *UserService) Profile(_ context.Context, id string) (*UserProfileItem, error) {
	user, err := dao.GetUserBasicInfo(id)
	if err != nil {
		logger.Errorw("GetUserBasicInfo", "err", err)
		return nil, err
	}
	userProfile, err := dao.GetUserProfile(id)
	if err != nil {
		logger.Errorw("GetUserProfile", "err", err)
		return nil, err
	}

	item := mergeUserProfile(user, userProfile)
	return item, nil
}

func mergeUserProfile(user model.UserBasicInfo, profile model.UserProfileInfo) *UserProfileItem {
	item := &UserProfileItem{
		UserID:     strconv.Itoa(user.ID),
		Username:   user.Username,
		Mobile:     user.Mobile,
		Role:       user.Role,
		State:      user.State,
		AuthStatus: profile.AuthStatus,

		AuthApplyTs: profile.AuthApplyAt.Unix(),
		RealName:    profile.RealName,
		Avatar:      profile.Avatar,
		Mail:        profile.Email,
		IDCard:      profile.IdCard,
		Gender:      profile.Gender,
		Hobby:       profile.Hobby,
		Career:      profile.Career,
		IDCardFront: profile.IdCardFront,
		IDCardBack:  profile.IdCardBack,
		CreateTs:    profile.CreatedAt.Unix(),
	}
	return item
}

func ProfileUpdate(_ context.Context, id string, params ProfileUpdateRequest) (*UserProfileItem, error) {
	// TODO
	return nil, nil
}

func (s *UserService) PublishOrUpdateRentalDemand(_ context.Context, id int, req RentalDemandRequest) error {
	info, modelID := formatRentalDemandParams(id, req)
	if err := dao.CreateOrUpdateUserRentalDemand(info, modelID); err != nil {
		logger.Errorw("AddUserRentalDemand", "err", err)
		return err
	}
	return nil
}

func formatRentalDemandParams(userID int, rentalDemand RentalDemandRequest) (*model.UserRentalDemandInfo, int) {
	info := &model.UserRentalDemandInfo{
		// Id:            rentalDemand.ID,
		UserID:        userID,
		DemandTitle:   rentalDemand.DemandTitle,
		ExtendContent: rentalDemand.ExtendContent,
		City:          rentalDemand.City,
		// RentTypeList:         "",
		// HouseTypeList:        "",
		// HouseFacilities:      "",
		// TrafficInfoJson: rentalDemand.TrafficInfoJson,
		MinMoneyBudget: rentalDemand.MinMoneyBudget,
		MaxMoneyBudget: rentalDemand.MaxMoneyBudget,
		Lighting:       rentalDemand.Lighting,
		// Floors:               "",
		CommutingTime:        rentalDemand.CommutingTime,
		CompanyAddress:       rentalDemand.CompanyAddress,
		DesiredResidenceArea: rentalDemand.DesiredResidenceArea,
		Elevator:             int(rentalDemand.Elevator),
		// State:                ,
		// JsonExtend:           nil,
		// CreatedAt:            time.Time{},
		// UpdatedAt:            time.Time{},
	}

	// 将整数列表转换为字符串列表，然后用 '#' 连接成一个字符串
	houseFacilities := make([]string, len(rentalDemand.HouseFacilities))
	for i, v := range rentalDemand.HouseFacilities {
		houseFacilities[i] = strconv.Itoa(v)
	}
	info.HouseFacilities = strings.Join(houseFacilities, "#")

	floors := make([]string, len(rentalDemand.Floors))
	for i, v := range rentalDemand.Floors {
		floors[i] = strconv.Itoa(v)
	}
	info.Floors = strings.Join(floors, "#")

	// 将枚举列表转换为字符串列表，然后用 '#' 连接成一个字符串
	rentTypeList := make([]string, len(rentalDemand.RentTypeList))
	for i, v := range rentalDemand.RentTypeList {
		rentTypeList[i] = v
	}
	info.RentTypeList = strings.Join(rentTypeList, "#")

	houseTypeList := make([]string, len(rentalDemand.HouseTypeList))
	for i, v := range rentalDemand.HouseTypeList {
		houseTypeList[i] = v
	}
	info.HouseTypeList = strings.Join(houseTypeList, "#")

	return info, rentalDemand.ID
}
