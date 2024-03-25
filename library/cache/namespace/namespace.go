package namespace

import (
	"context"
	"time"

	"github.com/cuixiaojun001/linkhome/library/cache"
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

func (n *namespaced) Delete(ctx context.Context, key string) error {
	key = n.getNamespacedKey(key)
	return n.c.Delete(ctx, key)
}

func (n *namespaced) getNamespacedKey(key string) string {
	return n.prefix + key
}
