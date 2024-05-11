package qiniu_test

import (
	"github.com/cuixiaojun001/LinkHome/cmd/http/bootstrap"
	"github.com/cuixiaojun001/LinkHome/common/config"
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/library/unitest"
	"github.com/cuixiaojun001/LinkHome/third_party"
	"github.com/cuixiaojun001/LinkHome/third_party/qiniu"
	"io/ioutil"
	"testing"
)

func init() {
	_ = bootstrap.SetUp(unitest.GetDevelopConfigFile())
	// 初始化第三方服务
	if err := third_party.Init(config.GetStringMap("third_party")); err != nil {
		logger.Fatalw("init third party failed", "err:", err)
	}
}

func Test_MakePrivateURL(t *testing.T) {
	url := qiniu.Client.MakePrivateURL("")
	t.Log(url)
}

func Test_UploadFile(t *testing.T) {
	// 解析本地路径下的文件为字节数组
	data, err := ioutil.ReadFile("/Users/ke/Documents/icon/storage.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	qiniu.Client.UploadFile(data)
}
