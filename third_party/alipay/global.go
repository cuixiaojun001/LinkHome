package alipay

//
//import (
//	"errors"
//	"fmt"
//	"github.com/cuixiaojun001/LinkHome/common/logger"
//	"github.com/cuixiaojun001/LinkHome/library/unitest"
//	"github.com/smartwalle/alipay"
//	"io"
//	"os"
//)
//
//var Client *AliPayClient
//
//func Init(config map[string]interface{}) error {
//	// appID, ok := config["appid"].(string)
//	if !ok {
//		return errors.New("[config] sms miss account parameter")
//	}
//	// 公钥和私钥从文件中获取
//	rootDir := unitest.GetCodeBasePath()
//	publicPath := rootDir + "/conf/keys/alipay_public_key.pem"
//	privatePath := rootDir + "/conf/keys/alipay_private_key.pem"
//	publicFile, err := os.Open(publicPath)
//	if err != nil {
//		logger.Fatalw("无法打开文件", "file:", publicPath)
//	}
//	defer publicFile.Close()
//	privateFile, err := os.Open(privatePath)
//	if err != nil {
//		logger.Fatalw("无法打开文件", "file:", privatePath)
//	}
//	defer privateFile.Close()
//	publicKey, err := io.ReadAll(publicFile)
//	if err != nil {
//		logger.Fatalw("无法读取文件", "file:"+publicPath)
//	}
//	privateKey, err := io.ReadAll(privateFile)
//	if err != nil {
//		logger.Fatalw("无法读取文件", "file:"+privatePath)
//	}
//	fmt.Println(string(publicKey))
//	fmt.Println(string(privateKey))
//
//	// alipay := alipay.New(appID, publicKey, privateKey, false)
//
//	// Client = New(alipay)
//	return nil
//}
