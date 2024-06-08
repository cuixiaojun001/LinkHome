package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
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

	// SAdd 将一个或多个成员元素加入到集合key当中，已经存在于集合的成员元素将被忽略。
	SAdd(ctx context.Context, key string, members ...interface{}) (int64, error)
	// SRem 将一个或多个成员元素从集合key当中移除，不存在的成员元素将被忽略。
	SRem(ctx context.Context, key string, members ...interface{}) (int64, error)
	// SMembers 返回集合key中的所有成员。
	SMembers(ctx context.Context, key string) ([]string, error)

	// ZAdd 将一个或多个成员元素及其分数值加入到有序集key当中。
	ZAdd(ctx context.Context, key string, members ...*redis.Z) (int64, error)
	// ZIncrBy 为有序集key的成员member的score值加上增量increment。 当 key 不存在，或分数不是 key 的成员时， ZINCRBY key increment member 等同于 ZADD key increment member 。
	ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error)
	// ZRevrRange 返回有序集key中，指定区间内的成员。其中成员的位置按score值递减(从大到小)来排列。
	ZRevrRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	// ZRevrRangeWithScores 返回有序集key中，指定区间内的成员。其中成员的位置按score值递减(从大到小)来排列。
	ZRevrRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error)

	// HIncrBy 为哈希表 key 中的域 field 的值加上增量 increment 。
	HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error)
	// HGetAll 返回哈希表 key 中，所有的域和值。
	HGetAll(ctx context.Context, key string) (map[string]string, error)

	// Keys 查找所有符合给定模式 pattern 的 key 。
	Keys(ctx context.Context, pattern string) ([]string, error)
	// Delete 从缓存中删除一个键。如果键存在，那么它将被删除，否则返回一个错误。
	Delete(ctx context.Context, key string) error
}
