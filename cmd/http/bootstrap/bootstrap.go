package bootstrap

import (
	"github.com/cuixiaojun001/LinkHome/common/cache"
	"github.com/cuixiaojun001/LinkHome/common/config"
	"github.com/cuixiaojun001/LinkHome/common/logger"
	"github.com/cuixiaojun001/LinkHome/common/mysql"
)

func SetUp(configFile string) error {
	// 初始化配置
	config.Init(configFile)

	// 初始化日志
	logger.SetUp()

	// 初始化DB
	mysql.SetUp()

	// 初始化缓存
	if err := cache.Init(config.GetStringMap("redis")); err != nil {
		return err
	}

	return nil
}

// Destroy 项目销毁
func Destroy() {
	mysql.DestroyMySQL()
}
