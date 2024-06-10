package order

import (
	"context"
	"fmt"
	"github.com/cuixiaojun001/LinkHome/third_party/qiniu"
	"strconv"
	"sync"
	"time"

	"github.com/cuixiaojun001/LinkHome/common/errdef"
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/library/orm"
	houseDao "github.com/cuixiaojun001/LinkHome/modules/house/dao"
	"github.com/cuixiaojun001/LinkHome/modules/order/dao"
	"github.com/cuixiaojun001/LinkHome/modules/order/model"
	userDao "github.com/cuixiaojun001/LinkHome/modules/user/dao"
	"github.com/cuixiaojun001/LinkHome/services/common"
)

type IOrderService interface {
	// CreateOrder 创建租房预定订单
	CreateOrder(ctx context.Context, userID int, request CreateOrderRequest) error
	// GetUserOrderList 获取用户订单列表
	GetUserOrderList(ctx context.Context, userID int) (*UserOrderListResponse, error)
}

type OrderService struct{}

var once sync.Once
var houseManager IOrderService

func GetOrderManager() IOrderService {
	once.Do(func() {
		houseManager = &OrderService{}
	})
	return houseManager
}

var _ IOrderService = (*OrderService)(nil)

func (s *OrderService) CreateOrder(ctx context.Context, userID int, request CreateOrderRequest) error {
	// 创建租房预定订单逻辑
	loc, _ := time.LoadLocation("Asia/Shanghai")
	startDate, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		logger.Debugw("time.Parse", "request.StartDate", request.StartDate)
		return err
	}
	startDate = startDate.Add(-8 * time.Hour).In(loc)
	endDate, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		logger.Debugw("time.Parse", "request.EndDate", request.EndDate)
		return err
	}
	endDate = endDate.Add(-8 * time.Hour).In(loc)
	// 判断入住和退租日期是否合理
	today := time.Now().In(loc)
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, loc)
	// logger.Debugw("CreateOrder", "today", today, "startOfDay", startOfDay, "startDate", startDate, "endDate", endDate)
	if startDate.Before(startOfDay) || endDate.Before(startOfDay) || startDate.After(startOfDay) {
		return errdef.DATE_ERR
	}
	// 获取房源拥有者
	houseDetail, err := houseDao.GetHouseDetail(request.HouseID)
	if err != nil {
		return err
	}
	// 同一租客id和同房源id的订单只有在当前订单处于结束状态才可以继续创建，避免重复创建
	notAllowedStatus := []interface{}{model.NoPay, model.Payed, model.Ordered, model.Deleted, model.Canceled}
	filter := orm.NewQuery().
		ExactMatch(model.TenantID, userID).
		ExactMatch(model.HouseID, houseDetail.HouseID).
		SetIn(model.State, notAllowedStatus)

	if exist, err := dao.CheckRecordExists(filter); err != nil {
		logger.Errorw("order FilterFirst", "err", err)
		return err
	} else if exist {
		return errdef.ORDER_EXIST_ERR
	}

	// 计算支付总金额 => 租金 * 租金扣押比率 + 租金 * 支付比率 + 管理费
	houseInfo, err := houseDao.GetHouseInfo(houseDetail.HouseID)
	payMoney := houseInfo.RentMoney*houseInfo.DepositRatio + houseInfo.RentMoney*houseInfo.PayRatio + houseInfo.StrataFee
	// 押金 => 租金 * 租金扣押比率
	depositFee := houseInfo.RentMoney * houseInfo.DepositRatio
	// 房屋预定金
	bargainMoney := houseInfo.BargainMoney
	// 创建订单
	rentalDays := endDate.Sub(startDate).Hours() / 24
	orderData := &model.OrderModel{
		// ID:              0,
		// TradeNo:         "",
		TenantID:   userID,
		LandlordID: houseDetail.HouseOwner,
		HouseID:    houseDetail.HouseID,
		// ContractContent: "",
		State:        model.NoPay,
		PayMoney:     payMoney,
		DepositFee:   depositFee,
		BargainMoney: bargainMoney,
		RentalDays:   int(rentalDays),
		StartDate:    startDate,
		EndDate:      endDate,
		//JsonExtend:   nil,
		//CreatedAt:    time.Time{},
		//UpdatedAt:    time.Time{},
	}
	if err := dao.CreateOrder(orderData); err != nil {
		logger.Errorw("CreateOrder", "err", err)
	}
	contract, err := common.GenerateContractContent(ctx, orderData.ID, 1)
	if err != nil {
		logger.Errorw("GenerateContractContent", "err", err)
		return err
	}
	orderData.ContractContent = contract
	if err := dao.UpdateOrder(orderData); err != nil {
		logger.Errorw("UpdateOrder", "err", err)
		return err
	}

	return nil
}

func (s *OrderService) GetUserOrderList(_ context.Context, userID int) (*UserOrderListResponse, error) {
	orders, err := dao.GetUserOrderList(userID)
	if err != nil {
		logger.Errorw("GetUserOrderList", "err", err)
	}
	var items []UserOrderListItem
	for _, order := range orders {
		houseInfo, err := houseDao.GetHouseInfo(order.HouseID)
		if err != nil {
			logger.Errorw("GetHouseInfo", "err", err)
			return nil, err
		}
		houseInfoItem := HouseInfoItem{
			HouseID:      houseInfo.ID,
			Title:        houseInfo.Title,
			Address:      houseInfo.Address,
			IndexImg:     qiniu.Client.MakePrivateURL(houseInfo.IndexImg),
			RentType:     houseInfo.RentType,
			RentMoney:    strconv.Itoa(houseInfo.RentMoney),
			StrataFee:    strconv.Itoa(houseInfo.StrataFee),
			DepositRatio: strconv.Itoa(houseInfo.DepositRatio),
			PayRatio:     strconv.Itoa(houseInfo.PayRatio),
		}
		userInfo, err := userDao.GetUserProfile(order.TenantID)
		if err != nil {
			logger.Errorw("GetUserProfile", "err", err)
			return nil, err
		}
		userInfoItem := UserInfoItem{
			UserID:   userInfo.Id,
			RearName: userInfo.RealName,
			Mobile:   userInfo.Mobile,
		}
		landlordInfo, err := userDao.GetUserProfile(order.LandlordID)
		if err != nil {
			logger.Errorw("GetUserProfile", "err", err)
			return nil, err
		}
		landlordInfoItem := UserInfoItem{
			UserID:   landlordInfo.Id,
			RearName: landlordInfo.RealName,
			Mobile:   landlordInfo.Mobile,
		}
		item := UserOrderListItem{
			OrderID:         order.ID,
			TenantID:        order.TenantID,
			LandlordID:      order.LandlordID,
			HouseID:         order.HouseID,
			StartDate:       order.StartDate.Format("2006-01-02"),
			EndDate:         order.EndDate.Format("2006-01-02"),
			State:           order.State,
			ContractContent: order.ContractContent,
			PayMoney:        strconv.Itoa(order.PayMoney),
			BargainMoney:    fmt.Sprintf("%f", order.BargainMoney),
			DepositFee:      strconv.Itoa(order.DepositFee),
			RentalDays:      order.RentalDays,
			UserInfo:        userInfoItem,
			LandlordInfo:    landlordInfoItem,
			HouseInfo:       houseInfoItem,
			CreateTs:        order.CreatedAt.Unix(),
			UpdateTs:        order.UpdatedAt.Unix(),
		}
		items = append(items, item)
	}

	return &UserOrderListResponse{UserOrders: items}, nil
}
