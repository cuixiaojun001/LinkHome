package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"

	"github.com/cuixiaojun001/LinkHome/library/cache"
	"github.com/go-redis/redis/v8"
)

const DriverName = "cache"

func init() {
	cache.Register(DriverName, &factory{})
}

type driver struct {
	client *redis.Client
}

func (d *driver) Set(ctx context.Context, key string, value interface{}) error {
	return d.setEx(ctx, key, value, 0)
}

func (d *driver) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return d.setEx(ctx, key, value, expiration)
}

func (d *driver) Get(ctx context.Context, key string, value interface{}) (bool, error) {
	result := d.client.Get(ctx, key)
	return d.loadValue(result, value)
}

func (d *driver) GetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	result := d.client.GetEx(ctx, key, expiration)
	return d.loadValue(result, value)
}

func (d *driver) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	result := d.client.SetNX(ctx, key, value, expiration)
	return result.Val(), result.Err()
}

func (d *driver) ZAdd(ctx context.Context, key string, members ...*redis.Z) (int64, error) {
	return d.client.ZAdd(ctx, key, members...).Result()
}

func (d *driver) ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	return d.client.ZIncrBy(ctx, key, increment, member).Result()
}

func (d *driver) ZRevrRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return d.client.ZRevRange(ctx, key, start, stop).Result()
}

func (d *driver) ZRevrRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	return d.client.ZRevRangeWithScores(ctx, key, start, stop).Result()
}

func (d *driver) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	return d.client.HIncrBy(ctx, key, field, incr).Result()
}

func (d *driver) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return d.client.HGetAll(ctx, key).Result()
}

func (d *driver) Keys(ctx context.Context, pattern string) ([]string, error) {
	return d.client.Keys(ctx, pattern).Result()
}

func (d *driver) TTL(ctx context.Context, key string) (int64, bool, error) {
	result := d.client.TTL(ctx, key)
	dur, err := result.Result()
	if err != nil {
		return 0, false, err
	}
	switch dur {
	case -2: // -2 if the key does not exist
		return 0, false, nil
	case -1: // -1 if the key exists but has no associated expire
		return math.MaxInt64, true, nil
	default:
		return time.Now().Add(dur).UnixNano(), true, nil
	}
}

func (d *driver) Delete(ctx context.Context, key string) error {
	return d.client.Del(ctx, key).Err()
}

func (d *driver) flush(ctx context.Context) error {
	result := d.client.FlushDB(ctx)
	if err := result.Err(); err != nil {
		return err
	}
	return nil
}

func (d *driver) setEx(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	var v interface{}

	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, []byte:
		v = value
	default:
		data, err := json.Marshal(value) // 将value转换为JSON格式的字节切片
		if err != nil {
			return err
		}
		v = data
	}

	return d.client.Set(ctx, key, v, expiration).Err()
}

func (d *driver) loadValue(result *redis.StringCmd, value interface{}) (exist bool, err error) {
	if err = result.Err(); err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}
	exist = true

	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		k := reflect.TypeOf(value)
		if k == nil {
			return exist, fmt.Errorf("cache: Get(nil)")
		}
		if k.Kind() != reflect.Ptr {
			return exist, fmt.Errorf("cache: Get(non-pointer " + k.String() + ")")
		}
		return exist, fmt.Errorf("cache: Get(nil " + k.String() + ")")
	}

	elem := rv.Elem()
	switch elem.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		data, err := result.Int64()
		elem.SetInt(data)
		return exist, err
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		data, err := result.Uint64()
		elem.SetUint(data)
		return exist, err
	case reflect.Float32, reflect.Float64:
		data, err := result.Float64()
		elem.SetFloat(data)
		return exist, err
	case reflect.Slice:
		if t := reflect.TypeOf(value).Elem(); t.Elem().Kind() == reflect.Uint8 {
			data, err := result.Bytes()
			if err != nil {
				return exist, err
			}
			elem.SetBytes(data)
			return exist, nil
		}
	case reflect.String:
		data := result.Val()
		elem.SetString(data)
		return exist, nil
	}

	data, err := result.Bytes()
	if err != nil {
		return exist, err
	}

	err = json.Unmarshal(data, value)
	if err != nil {
		return exist, err
	}

	return true, nil
}

func newDriver(cfg map[string]interface{}) (*driver, error) {
	addr, ok := cfg["addr"].(string)
	if !ok || len(addr) == 0 {
		return nil, errors.New("cache miss addr parameter")
	}

	var username, password string
	if c := cfg["username"]; c != nil {
		username, _ = c.(string)
	}
	if c := cfg["password"]; c != nil {
		password, _ = c.(string)
	}

	var database, poolSize int
	if c := cfg["database"]; c != nil {
		if database, ok = c.(int); !ok {
			return nil, errors.New("cache invalid database parameter, need int")
		}
	}
	if c := cfg["pool_size"]; c != nil {
		if poolSize, ok = c.(int); !ok {
			return nil, errors.New("cache invalid pool_size parameter, need int")
		}
	} else {
		poolSize = 64
	}

	cli := redis.NewClient(&redis.Options{
		Addr:         addr,
		Username:     username,
		Password:     password,
		DB:           database,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		PoolSize:     poolSize,
		MinIdleConns: 8,
	})

	if err := cli.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	client := &driver{client: cli}

	return client, nil
}

type factory struct{}

func (f *factory) New(cfg map[string]interface{}) (cache.Cache, error) {
	return newDriver(cfg)
}
