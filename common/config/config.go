package config

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/cuixiaojun001/linkhome/common/env"
	"github.com/shima-park/agollo"
	"github.com/spf13/viper"
)

var viperConfig *viper.Viper   // viperConfig viper客户端
var apolloClient agollo.Agollo // apollo 客户端

// Init 初始化配置
func Init(configFile string) {
	viperConfig = viper.New()
	viperConfig.SetConfigFile(configFile)
	err := viperConfig.ReadInConfig()
	if err != nil {
		log.Fatalf("配置文件加载失败: %v\n", err)
	}

	env.Env = env.EnvTypeOf(GetStringMust("environment"))

	//remote.SetAppID(viperConfig.GetString("apollo.service_id"))
	//remote.SetAgolloOptions(
	//	agollo.AutoFetchOnCacheMiss(),
	//	agollo.FailTolerantOnBackupExists(),
	//	agollo.BackupFile(".apollo"),
	//)
	//
	//viperConfig.AddRemoteProvider("apollo", viperConfig.GetString("apollo.endpoint"), viperConfig.GetString("apollo.bauth_namespace"))
	//if !(viperConfig.ReadRemoteConfig() == nil && viperConfig.WatchRemoteConfigOnChannel() == nil) {
	//	log.Fatal("配置中心加载失败\n")
	//}
	//
	//if err := initApolloClient(viperConfig); err != nil {
	//	log.Fatalf("init apollo client fail, err:%+v\n", err)
	//}
}

// 初始化 apollo 客户端
func initApolloClient(conf *viper.Viper) error {
	client, err := agollo.New(conf.GetString("apollo.endpoint"), conf.GetString("apollo.service_id"),
		agollo.AutoFetchOnCacheMiss(),
		agollo.EnableSLB(true),
		agollo.BackupFile(".agollo"),
		agollo.FailTolerantOnBackupExists(),
		agollo.LongPollerInterval(10*time.Second),
	)
	if err != nil {
		return err
	}

	client.Start()
	apolloClient = client

	return nil
}

// IsSet 判断配置项是否存在
func IsSet(key string) bool {
	return viperConfig.IsSet(key)
}

// AllSettings 获取所有的配置信息
func AllSettings() map[string]interface{} {
	return viperConfig.AllSettings()
}

// GetStringMap 根据name获取配置信息
func GetStringMap(name string) map[string]interface{} {
	return viperConfig.GetStringMap(name)
}

// GetStringSlice 根据name获取配置信息
func GetStringSlice(name string) []string {
	return viperConfig.GetStringSlice(name)
}

// GetString 根据name获取配置项的值
func GetString(name string) (string, error) {
	if viperConfig.IsSet(name) {
		return viperConfig.GetString(name), nil
	}
	return "", errors.New("配置项不存在: " + name)
}

// GetInt 根据name获取配置项的整数值
func GetInt(name string, base int, bitSize int) (int64, error) {
	v, err := GetString(name)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(v, base, bitSize)
}

// GetUint 根据name获取配置项的无符号整数值
func GetUint(name string, base int, bitSize int) (uint64, error) {
	v, err := GetString(name)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(v, base, bitSize)
}

// GetFloat 根据name获取配置项的小数值
func GetFloat(name string, bitSize int) (float64, error) {
	v, err := GetString(name)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(v, bitSize)
}

// GetBool 根据name获取配置项的布尔值
func GetBool(name string) (bool, error) {
	v, err := GetString(name)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(v)
}

// GetStringMust 根据name获取配置项的值，没有会退出
func GetStringMust(name string) string {
	v, err := GetString(name)
	if err != nil {
		log.Fatal("Get configuration fails", "config name:", name, "err:", err)
	}
	return v
}

// GetIntMust 根据name获取配置项的整数值，没有会退出
func GetIntMust(name string, base int, bitSize int) int64 {
	v, err := GetInt(name, base, bitSize)
	if err != nil {
		log.Fatal("Get configuration fails", "config name:", name, "err:", err)
	}
	return v
}

// GetBoolMust 根据name获取配置项的布尔值，没有会退出
func GetBoolMust(name string) bool {
	v, err := GetBool(name)
	if err != nil {
		log.Fatal("Get configuration fails", "config name:", name, "err:", err)
	}
	return v
}

// Unmarshal 将配置信息解析到对应的数据结构
func Unmarshal(i interface{}, opt ...viper.DecoderConfigOption) error {
	return viperConfig.Unmarshal(i, opt...)
}

func UnmarshalKey(key string, i interface{}, opt ...viper.DecoderConfigOption) error {
	return viperConfig.UnmarshalKey(key, i, opt...)
}

func Config() *viper.Viper {
	return viperConfig
}

// Sub 根据key获取相应的配置结构
func Sub(key string) *viper.Viper {
	return viperConfig.Sub(key)
}
