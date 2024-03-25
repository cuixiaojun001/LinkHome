package user

import (
	"context"
	"errors"
	"github.com/cuixiaojun001/linkhome/modules/user/dao"
	"github.com/cuixiaojun001/linkhome/modules/user/model"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"time"

	"github.com/cuixiaojun001/linkhome/common/errdef"
	"github.com/cuixiaojun001/linkhome/common/logger"
	"github.com/cuixiaojun001/linkhome/third_party/sms"
	"github.com/golang-jwt/jwt"
)

const (
	PhoneRegex           = `^1[3456789]\d{9}$`
	JwtSecret            = "NmUzODk2ZGUtYmZjYy0xMWVjLWI5YTctZjQzMGI5YTUwMzQ2aHVp"
	JwtExpiryHours       = 2
	JwtRefreshExpiryDays = 14
)

func Login(_ context.Context, req LoginRequest) (*TokenItem, error) {
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

func IsUsernameExist(username string) bool {
	return dao.IsUsernameExist(username)
}

func IsMobileExist(mobile string) bool {
	return dao.IsMobileExist(mobile)
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

func SendSmsCode(_ context.Context, mobile string) error {
	errCh := make(chan error)
	go sms.Client.SendSmsCode(mobile, errCh)

	err := <-errCh
	if err != nil {
		logger.Errorw("SendSmsCode", "err", err)
		return err
	}
	return nil
}

func Register(_ context.Context, req RegisterRequest) (*TokenItem, error) {
	if err := verifyUserRegisterInfo(req); err != nil {
		logger.Infow("verifyUserRegisterInfo", "err", err)
		return nil, err
	}
	// 注册用户
	user := model.UserBasicInfo{
		Username: req.Username,
		Password: req.Password,
		Role:     string(req.Role),
		Mobile:   req.Mobile,
		State:    "normal",
	}
	if err := dao.AddUser(&user); err != nil {
		logger.Errorw("AddUser", "err", err)
		return nil, err
	}
	// TODO 向用户详情表添加用户记录

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

func PwdChange(_ context.Context, userID string, req PwdChangeRequest) (*TokenItem, error) {
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

func Profile(_ context.Context, id string) (*UserProfileItem, error) {
	user, err := dao.GetUserBasicInfoByID(id)
	if err != nil {
		logger.Errorw("GetUserBasicInfo", "err", err)
		return nil, err
	}
	userProfile, err := dao.GetUserProfileByID(id)
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
