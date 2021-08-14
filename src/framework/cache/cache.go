package cache

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mixnote/mixnote-api-go/src/framework/utilities"
)

var (
	once   sync.Once
	ctx    = context.Background()
	_redis *redis.Client
)

func NewCache(db int) {
	once.Do(func() {
		_redis = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       db,
		})
		if res := _redis.Ping(ctx); res.Err() != nil {
			utilities.Console.Fatal(res.Err())
		}
	})
}

func Set(key string, val interface{}, expires time.Duration) *redis.StatusCmd {
	defer ctx.Done()
	return _redis.Set(ctx, key, val, expires)
}

func Get(key string) (string, error) {
	defer ctx.Done()
	return _redis.Get(ctx, key).Result()
}

func Exist(key string) bool {
	defer ctx.Done()
	_, err := _redis.Get(ctx, key).Result()
	if err == redis.Nil{
		return false
	} else if err != nil {
		utilities.Console.Fatal(err)
	}
	return true
}

func Forget() {}

func Forever() {}

func UseDB(db int) {
	if _redis != nil {
		_redis.Close()
	}

	_redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       db,
	})
	if res := _redis.Ping(ctx); res.Err() != nil {
		utilities.Console.Fatal(res.Err())
	}
}
