package sms

import (
	"fmt"
	"time"
)

// AppName 应用名称
const AppName = "LinkHome"

// CodeTimeout 验证码的超时时间
const CodeTimeout = 5 * time.Minute

// RedisDataTypeString Redis数据类型为字符串
const RedisDataTypeString = "string"

// RedisCacheInfo 定义一个结构体，表示Redis缓存信息
type RedisCacheInfo struct {
	Key      string
	Timeout  time.Duration
	DataType string
}

// mobileSMSCode 用于获取手机验证码的Redis缓存信息
func mobileSMSCode(mobile string) RedisCacheInfo {
	smsCodeCacheInfo := RedisCacheInfo{
		Key:      fmt.Sprintf("%s:user:sms_code:%s", AppName, mobile),
		Timeout:  CodeTimeout,
		DataType: RedisDataTypeString,
	}

	return smsCodeCacheInfo
}
