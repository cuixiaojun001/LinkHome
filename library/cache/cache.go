package cache

import (
	"context"
	"time"
)

// Cache 缓存接口
type Cache interface {
	// Set 设置一个键值对(永不过期)。如果键已经存在，那么它的值将被更新。
	Set(ctx context.Context, key string, value interface{}) error
	// SetEX 设置一个键值对，并指定一个过期时间。如果键已经存在，那么它的值和过期时间将被更新。
	SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error

	// Get 从缓存中获取一个键的值。如果键存在，那么它的值将被返回，否则返回一个错误。
	Get(ctx context.Context, key string, value interface{}) (bool, error)
	// GetEX 从缓存中获取一个键的值，并更新它的过期时间。如果键存在，那么它的值将被返回，否则返回一个错误。
	GetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	// SetNX 在缓存中设置一个不存在的键值对，并指定一个过期时间。如果键已经存在，那么这个操作将失败。
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)

	// TTL 获取一个键的剩余生存时间。如果键存在，那么它的剩余生存时间将被返回，否则返回一个错误。
	TTL(ctx context.Context, key string) (int64, bool, error)

	// Delete 从缓存中删除一个键。如果键存在，那么它将被删除，否则返回一个错误。
	Delete(ctx context.Context, key string) error
}
