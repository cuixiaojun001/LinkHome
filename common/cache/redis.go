package cache

import (
	"github.com/cuixiaojun001/linkhome/library/cache"
	"github.com/cuixiaojun001/linkhome/library/cache/namespace"
	redisCache "github.com/cuixiaojun001/linkhome/library/cache/redis"
)

type Cache = cache.Cache

// TODO 后期实现Tiered Cache分层缓存
var rediscache Cache

func Init(cfg map[string]interface{}) (err error) {
	rediscache, err = cache.New(redisCache.DriverName, cfg)
	if err != nil {
		return err
	}

	return nil
}

// New 创建一个 cache 客户端
// 参数 ns 表示所属于空间，例如: New("application") 创建一个 application 使用的缓存客户端
// New("user-center") 创建一个 user-center 使用的缓存客户端
func New(ns string) Cache {
	return namespace.Namespaced(ns, rediscache)
}
