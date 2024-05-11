package sms_test

import (
	"testing"

	"github.com/cuixiaojun001/LinkHome/cmd/http/bootstrap"
	"github.com/cuixiaojun001/LinkHome/common/config"
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/library/unitest"
	"github.com/cuixiaojun001/LinkHome/third_party"
)

func init() {
	_ = bootstrap.SetUp(unitest.GetDevelopConfigFile())
	// 初始化第三方服务
	if err := third_party.Init(config.GetStringMap("third_party")); err != nil {
		logger.Fatalw("init third party failed", "err:", err)
	}
}

func TestSendSmsCode(t *testing.T) {
	//if err := sms.Client.SendSmsCode("", ); err != nil {
	//	t.Log(err)
	//}
}
