package env

import (
	"log"
	"strings"
)

// Env 当前应用运行环境
var Env string

const (
	Develop = "dev"     // 开发环境
	Test    = "test"    // 测试环境
	Preview = "preview" // 预览环境
	Product = "prod"    // 生产环境
)

// EnvTypeOf 字符串转换为常量
func EnvTypeOf(envType string) string {
	envType = strings.ToLower(envType)
	switch envType {
	case "dev", "develop":
		return Develop
	case "test":
		return Test
	case "preview":
		return Preview
	case "prod", "product":
		return Product
	default:
		log.Fatal("envType值存在问题，必须为下面几个值之一：dev, develop, test, preview, prod, product")
		return Develop
	}
}

// EnvCnOf 字符串转中文
func EnvCnOf(envType string) string {
	envType = EnvTypeOf(envType)
	switch envType {
	case Develop:
		return "开发环境"
	case Test:
		return "测试环境"
	case Preview:
		return "预发布环境"
	case Product:
		return "生产环境"
	default:
		return "开发环境"
	}
}
