package alipay

import (
	"errors"
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/library/unitest"
	"github.com/smartwalle/alipay/v3"
	"log"
)

var Client *AliPayClient
var client *alipay.Client

func Init(config map[string]interface{}) (err error) {
	appID, ok := config["appid"].(string)
	if !ok {
		return errors.New("[config] sms miss account parameter")
	}
	appPrivateKey, ok := config["app_private_key"].(string)
	if !ok {
		return errors.New("[config] sms miss account parameter")
	}
	callbackURL, ok := config["callback_url"].(string)
	if !ok {
		return errors.New("[config] sms miss account parameter")
	}

	if client, err = alipay.New(appID, appPrivateKey, false); err != nil {
		logger.Errorw("初始化支付宝失败", "err:", err)
		return err
	}

	rootDir := unitest.GetCodeBasePath()
	// 加载证书
	if err = client.LoadAppCertPublicKeyFromFile(rootDir + "/conf/cert/appPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return err
	}
	if err = client.LoadAliPayRootCertFromFile(rootDir + "/conf/cert/alipayRootCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return err
	}
	if err = client.LoadAlipayCertPublicKeyFromFile(rootDir + "/conf/cert/alipayPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return err
	}

	if err = client.SetEncryptKey("SYwfdwXto+Dt88EcLSyM/A=="); err != nil {
		log.Println("加载内容加密密钥发生错误", err)
		return
	}

	Client = New(client, callbackURL)

	return nil
}
