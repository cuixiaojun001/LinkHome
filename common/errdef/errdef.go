package errdef

import (
	"fmt"
)

var (
	OK    = &CustomError{0, "SUCCESS"}
	ERROR = &CustomError{-1, "FAILED"}

	IMAGE_CODE_ERR      = &CustomError{4001, "图形验证码错误"}
	THROTTLING_ERR      = &CustomError{4002, "访问过于频繁"}
	NECESSARY_PARAM_ERR = &CustomError{4003, "缺少必传参数"}
	AccountErr          = &CustomError{4004, "账号或密码错误"}
	AUTHORIZATION_ERR   = &CustomError{4005, "权限认证错误"}
	CpwdErr             = &CustomError{4006, "密码不一致"}
	MOBILE_ERR          = &CustomError{4007, "手机号错误"}
	SmsCodeErr          = &CustomError{4008, "短信验证码有误"}
	ALLOW_ERR           = &CustomError{4009, "未勾选协议"}
	SESSION_ERR         = &CustomError{4010, "用户未登录"}
	REGISTER_FAILED_ERR = &CustomError{4011, "注册失败"}
	FACILITY_EXIST_ERR  = &CustomError{4012, "房屋设施已存在"}
	PUBLISH_HOUSE_ERR   = &CustomError{4013, "发布房源失败"}
	DATE_ERR            = &CustomError{4014, "日期错误"}
	ORDER_EXIST_ERR     = &CustomError{4015, "订单已存在"}
	ORDER_INFO_ERR      = &CustomError{4016, "订单信息错误"}
	FORBIDDEN_ERR       = &CustomError{4017, "非法请求"}
	REALNAME_AUTH_ERR   = &CustomError{4018, "实名认证错误"}
	DB_ERR              = &CustomError{5000, "数据库错误"}
	EMAIL_ERR           = &CustomError{5001, "邮箱错误"}
	TEL_ERR             = &CustomError{5002, "固定电话错误"}
	NODATA_ERR          = &CustomError{5003, "无数据"}
	NEW_PWD_ERR         = &CustomError{5004, "新密码错误"}
	OPENID_ERR          = &CustomError{5005, "无效的openid"}
	PARAM_ERR           = &CustomError{5006, "参数错误"}
	STOCK_ERR           = &CustomError{5007, "库存不足"}
	SOCKET_ERR          = &CustomError{5008, "网络错误"}
	SYSTEM_ERR          = &CustomError{5009, "系统错误"}
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func New(code int, msg string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: msg,
	}
}
