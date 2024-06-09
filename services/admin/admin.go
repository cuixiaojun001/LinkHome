package admin

import (
	"context"
	"github.com/cuixiaojun001/LinkHome/services/house"
	"strconv"
	"sync"

	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/library/orm"
	"github.com/cuixiaojun001/LinkHome/modules/admin/dao"
	"github.com/cuixiaojun001/LinkHome/modules/user/model"
	"github.com/cuixiaojun001/LinkHome/services/order"
	"github.com/cuixiaojun001/LinkHome/services/user"
)

type IAdminService interface {
	// GetUserList 获取用户列表
	GetUserList(ctx context.Context, req *UserListRequest) (*UserListResponse, error)
	// GetOrderList 获取订单列表
	GetOrderList(ctx context.Context, req *OrderListRequest) (*OrderListResponse, error)
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

func (s *AdminService) GetOrderList(ctx context.Context, req *OrderListRequest) (*OrderListResponse, error) {
	filter := req.GenQuery()
	orderList, err := dao.GetOrder(filter)
	if err != nil {
		return nil, err
	}

	// 使用map来记录已经添加过的TenantID
	seen := make(map[int]bool)
	userIDs := make([]int, 0, len(orderList))

	for _, order := range orderList {
		if _, ok := seen[order.TenantID]; !ok {
			// 如果TenantID还没有被添加过，则添加到userIDs中
			userIDs = append(userIDs, order.TenantID)
			// 并标记为已添加
			seen[order.TenantID] = true
		}
	}

	var items []UserOrderListItem
	orderMgr := order.GetOrderManager()
	for _, id := range userIDs {
		resp, err := orderMgr.GetUserOrderList(ctx, id)
		if err != nil {
			logger.Errorw("GetUserOrderList failed", "err", err)
			continue
		}
		for _, item := range resp.UserOrders {
			houseMgr := house.GetHouseManager()
			house, err := houseMgr.GetHouseDetail(ctx, item.HouseID, item.LandlordID)
			if err != nil {
				logger.Errorw("GetHouseDetail failed", "err", err)
				continue
			}
			items = append(items, UserOrderListItem{
				OrderID:         item.OrderID,
				TenantID:        item.TenantID,
				LandlordID:      item.LandlordID,
				HouseID:         item.HouseID,
				StartDate:       item.StartDate,
				EndDate:         item.EndDate,
				State:           item.State,
				ContractContent: item.ContractContent,
				PayMoney:        item.PayMoney,
				BargainMoney:    item.BargainMoney,
				DepositFee:      item.DepositFee,
				RentalDays:      item.RentalDays,
				UserInfo:        item.UserInfo,
				LandlordInfo:    item.LandlordInfo,
				HouseInfo:       *house,
				CreateTs:        item.CreateTs,
				UpdateTs:        item.UpdateTs,
			})
		}
	}
	return &OrderListResponse{
		Total:      len(items),
		HasMore:    false,
		NextOffset: 0,
		DataList:   items,
	}, err
}
