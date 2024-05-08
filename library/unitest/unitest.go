package unitest

import (
	"path"
	"runtime"
)

// GetCodeBasePath 获取代码根目录
func GetCodeBasePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(path.Dir(path.Dir(filename)))
}

func GetDevelopConfigFile() string {
	return GetCodeBasePath() + "/conf/dev.yaml"
}
