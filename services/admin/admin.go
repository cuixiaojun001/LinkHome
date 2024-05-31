package admin

import (
	"context"
	"strconv"
	"sync"

	"github.com/cuixiaojun001/LinkHome/library/orm"
	"github.com/cuixiaojun001/LinkHome/modules/admin/dao"
	"github.com/cuixiaojun001/LinkHome/modules/user/model"
	"github.com/cuixiaojun001/LinkHome/services/user"
)

type IAdminService interface {
	// GetUserList 获取用户列表
	GetUserList(ctx context.Context, req *UserListRequest) (*UserListResponse, error)
}

type AdminService struct{}

var once sync.Once
var adminManager IAdminService

func GetAdminManager() IAdminService {
	once.Do(func() {
		adminManager = &AdminService{}
	})
	return adminManager
}

var _ IAdminService = (*AdminService)(nil)

func (s *AdminService) GetUserList(ctx context.Context, req *UserListRequest) (*UserListResponse, error) {
	filter := req.GenQuery()
	profiles, err := dao.GetUserProfile(filter)
	if err != nil {
		return nil, err
	}

	var ids []int
	temp := make(map[int]model.UserProfileInfo, len(profiles))
	for _, profile := range profiles {
		ids = append(ids, profile.Id)
		temp[profile.Id] = profile
	}

	filter = orm.NewQuery().SetInInt("id", ids)
	basics, err := dao.GetUserBasic(filter)
	if err != nil {
		return nil, err
	}

	var items []user.UserProfileItem
	for _, basic := range basics {
		profile := temp[basic.ID]
		item := mergeUserProfile(basic, profile)
		items = append(items, item)
	}

	return &UserListResponse{
		Total:      len(items),
		HasMore:    false,
		NextOffset: 0,
		DataList:   items,
	}, err
}

func mergeUserProfile(basic model.UserBasicInfo, profile model.UserProfileInfo) user.UserProfileItem {
	item := user.UserProfileItem{
		UserID:     strconv.Itoa(basic.ID),
		Username:   basic.Username,
		Mobile:     basic.Mobile,
		Role:       basic.Role,
		State:      basic.State,
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
