package api

import (
	"context"
	"github.com/cuixiaojun001/LinkHome/common/config"
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/common/response"
	"github.com/cuixiaojun001/LinkHome/services/payment"
	pay "github.com/cuixiaojun001/LinkHome/third_party/alipay"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"net/http"
	"strconv"
)

func AliPayOrder(c *gin.Context) {
	orderIDStr := c.Param("order_id")
	orderID, _ := strconv.Atoi(orderIDStr)

	req := &payment.OrderPaymentRequest{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}
	if resp, err := payment.AliPayOrder(orderID, req); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
	} else {
		c.JSON(http.StatusOK, response.Success(resp))
	}
}

func AliPayCallback(c *gin.Context) {
	c.Request.ParseForm()
	// 移除指定的查询参数 `pay_scene`
	payScene := c.Request.Form.Get("pay_scene")
	c.Request.Form.Del("pay_scene")
	if err := pay.Client.GetClient().VerifySign(c.Request.Form); err != nil {
		logger.Errorw("回调验证签名发生错误", "err:", err)
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}

	logger.Debugw("回调验证签名通过")
	outTradeNo, _ := strconv.Atoi(c.Request.Form.Get("out_trade_no"))
	tradeNo := c.Request.Form.Get("trade_no")
	totalAmount, _ := strconv.Atoi(c.Request.Form.Get("total_amount"))
	p := alipay.TradeQuery{
		OutTradeNo: strconv.Itoa(outTradeNo),
	}
	rsp, err := pay.Client.GetClient().TradeQuery(context.TODO(), p)
	if err != nil || rsp.IsFailure() {
		logger.Errorw("验证订单 %s 信息发生错误:", "订单：", outTradeNo, "err:", err)
		c.JSON(http.StatusOK, response.BadRequest(err))
		return
	}

	if err := payment.AliPayCallBack(context.Background(), outTradeNo, totalAmount, tradeNo, payScene); err != nil {
		logger.Errorw("保存支付流水失败")
	}

	// 发送重定向请求
	http.Redirect(c.Writer, c.Request, config.GetStringMust("domain")+"/order.html", http.StatusTemporaryRedirect)
}
