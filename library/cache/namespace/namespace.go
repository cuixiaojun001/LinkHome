package namespace

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/cuixiaojun001/LinkHome/library/cache"
)

func Namespaced(namespace string, c cache.Cache) cache.Cache {
	prefix := ""
	if namespace != "" {
		prefix = namespace + ":"
	}
	return &namespaced{
		c:      c,
		prefix: prefix,
	}
}

type namespaced struct {
	c      cache.Cache
	prefix string
}

func (n *namespaced) Set(ctx context.Context, key string, value interface{}) error {
	key = n.getNamespacedKey(key)
	return n.c.Set(ctx, key, value)
}

func (n *namespaced) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	key = n.getNamespacedKey(key)
	return n.c.SetEX(ctx, key, value, expiration)
}

func (n *namespaced) Get(ctx context.Context, key string, value interface{}) (bool, error) {
	key = n.getNamespacedKey(key)
	return n.c.Get(ctx, key, value)
}

func (n *namespaced) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	key = n.getNamespacedKey(key)
	return n.c.SetNX(ctx, key, value, expiration)
}

func (n *namespaced) GetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	key = n.getNamespacedKey(key)
	return n.c.GetEX(ctx, key, value, expiration)
}

func (n *namespaced) TTL(ctx context.Context, key string) (int64, bool, error) {
	key = n.getNamespacedKey(key)
	return n.c.TTL(ctx, key)
}

func (n *namespaced) SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {
	key = n.getNamespacedKey(key)
	return n.c.SAdd(ctx, key, members...)
}

func (n *namespaced) SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	key = n.getNamespacedKey(key)
	return n.c.SRem(ctx, key, members...)
}

func (n *namespaced) SMembers(ctx context.Context, key string) ([]string, error) {
	key = n.getNamespacedKey(key)
	return n.c.SMembers(ctx, key)
}

func (n *namespaced) ZAdd(ctx context.Context, key string, members ...*redis.Z) (int64, error) {
	key = n.getNamespacedKey(key)
	return n.c.ZAdd(ctx, key, members...)
}

func (n *namespaced) ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	key = n.getNamespacedKey(key)
	return n.c.ZIncrBy(ctx, key, increment, member)
}

func (n *namespaced) ZRevrRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	key = n.getNamespacedKey(key)
	return n.c.ZRevrRange(ctx, key, start, stop)
}

func (n *namespaced) ZRevrRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	key = n.getNamespacedKey(key)
	return n.c.ZRevrRangeWithScores(ctx, key, start, stop)
}

func (n *namespaced) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	key = n.getNamespacedKey(key)
	return n.c.HIncrBy(ctx, key, field, incr)
}

func (n *namespaced) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	key = n.getNamespacedKey(key)
	return n.c.HGetAll(ctx, key)
}

func (n *namespaced) Keys(ctx context.Context, pattern string) ([]string, error) {
	return n.c.Keys(ctx, pattern)
}

func (n *namespaced) Delete(ctx context.Context, key string) error {
	key = n.getNamespacedKey(key)
	return n.c.Delete(ctx, key)
}

func (n *namespaced) getNamespacedKey(key string) string {
	return n.prefix + key
}
