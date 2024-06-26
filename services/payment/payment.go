package payment

import (
	"context"
	"github.com/cuixiaojun001/LinkHome/common/logger"
	houseDao "github.com/cuixiaojun001/LinkHome/modules/house/dao"
	orderDao "github.com/cuixiaojun001/LinkHome/modules/order/dao"
	"github.com/cuixiaojun001/LinkHome/modules/payment/dao"
	"github.com/cuixiaojun001/LinkHome/modules/payment/model"
	pay "github.com/cuixiaojun001/LinkHome/third_party/alipay"
	"github.com/smartwalle/alipay/v3"
	"strconv"
)

func AliPayOrder(orderID int, req *OrderPaymentRequest) (*OrderPaymentResponse, error) {
	order, err := orderDao.GetOrder(orderID)
	if err != nil && order.State != "deleted" {
		logger.Errorw("AliPayOrder", "GetOrder", err)
		return nil, err
	}

	if req.PayScene == "balance_payment" && order.State == "ordered" {
		// 已预订支付余款，由于支付宝同一个订单号不能支付多次,先创建新的订单然后把旧的订单删掉, 这样订单号就不同了但数据一样
		order, err = orderDao.UpdateOrderID(order)
		if err != nil {
			logger.Errorw("AliPayOrder", "UpdateOrderID", err)
			return nil, err
		}
	}

	// 判断是否需要修改租房开始日期和结束日期
	var startDate, endDate string
	if req != nil {
		startDate = req.StartDate
		endDate = req.EndDate
	}
	if startDate != "" && endDate != "" && (startDate != order.StartDate.Format("2006-01-02") || endDate != order.EndDate.Format("2006-01-02")) {
		// TODO 更新订单表中起止日期
	}

	totalAmount := 0
	if req.PayScene == "full_payment" {
		totalAmount = order.PayMoney
	} else if req.PayScene == "deposit_payment" {
		totalAmount = int(order.BargainMoney)
	} else if req.PayScene == "balance_payment" && order.State == "ordered" {
		totalAmount = order.PayMoney - int(order.BargainMoney)
	}

	// 生成订单信息字符串（再回调时会自动返回）
	tradeNo := strconv.Itoa(orderID)
	p := alipay.TradePagePay{
		Trade: alipay.Trade{
			AuxParam:    alipay.AuxParam{},
			NotifyURL:   pay.Client.CallBackURL() + "/alipay/notify",
			ReturnURL:   pay.Client.CallBackURL() + "?pay_scene=" + req.PayScene,
			Subject:     "LinkHome:本次支付订单号:" + tradeNo,
			OutTradeNo:  tradeNo,
			TotalAmount: strconv.Itoa(totalAmount),
			ProductCode: "FAST_INSTANT_TRADE_PAY",
		},
	}

	url, _ := pay.Client.GetClient().TradePagePay(p)

	return &OrderPaymentResponse{
		OrderID:   orderID,
		AliPayURL: url.String(),
	}, nil
}

func AliPayCallBack(ctx context.Context, orderID, transAmount int, tradeNo, scene string) error {
	order, err := orderDao.GetOrder(orderID)
	if err != nil && order.State != "deleted" {
		logger.Errorw("AliPayOrder", "GetOrder", err)
		return err
	}
	houseInfo, err := houseDao.GetHouseInfo(order.HouseID)
	if err != nil {
		logger.Errorw("AliPayOrder", "GetHouseInfo", err)
		return err
	}

	if scene == "full_payment" {
		order.State = "payed"
		houseInfo.RentState = "rent"
	} else if scene == "balance_payment" {
		order.State = "payed"
		houseInfo.RentState = "rent"
	} else if scene == "deposit_payment" {
		order.State = "ordered"
		houseInfo.RentState = "ordered"
	}

	if err := orderDao.UpdateOrder(order); err != nil {
		logger.Errorw("AliPayCallBackOrder", "UpdateOrder", err)
		return err
	}

	if err = houseDao.UpdateHouseInfo(houseInfo); err != nil {
		logger.Errorw("AliPayOrder", "UpdateHouseInfo", err)
		return err
	}

	trade := &model.PaymentTrade{
		//ID:          0,
		OrderId:     order.ID,
		UserId:      order.TenantID,
		TradeNo:     tradeNo,
		Scene:       scene,
		TransAmount: transAmount,
	}

	if err = dao.CreatePaymentTrade(trade); err != nil {
		logger.Errorw("AliPayOrder", "CreatePaymentTrade", err)
		return err
	}

	return nil
}
