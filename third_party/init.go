package third_party

import (
	"errors"
	"fmt"
	"github.com/cuixiaojun001/LinkHome/third_party/qiniu"
	"github.com/cuixiaojun001/LinkHome/third_party/sms"
)

// Module 第三方模块
type Module struct {
	Name string
	Init func(map[string]interface{}) error
}

var modules []Module

func init() {
	modules = append(modules, Module{Name: "sms", Init: sms.Init})     // 短信服务
	modules = append(modules, Module{Name: "qiniu", Init: qiniu.Init}) // 七牛云对象存储服务
	// modules = append(modules, Module{Name: "alipay", Init: alipay.Init})
}

// Init 根据配置文件初始化所有以来的第三方模块
func Init(conf map[string]interface{}) error {
	for _, m := range modules {
		if config, ok := conf[m.Name].(map[string]interface{}); !ok {
			return errors.New(fmt.Sprintf("配置文件中缺少%s的配置", m.Name))
		} else if err := m.Init(config); err != nil {
			return err
		}
	}

	return nil
}
