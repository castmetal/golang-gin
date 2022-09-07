package redis

import (
	"context"
	"fmt"
	_config "golang-gin/src/config"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

type IRedisClient interface {
	GetData(ctx context.Context, key string) (string, error)
	SetData(ctx context.Context, key string, value string, ttl time.Duration) error
	DelData(ctx context.Context, keyPattern string) error
}

type RedisConn struct {
	Client *redis.Client
}

func NewRedisClient() IRedisClient {
	port, err := strconv.Atoi(_config.SystemParams.REDIS_PORT)
	if err != nil {
		port = 6379
	}

	redisDb, err := strconv.Atoi(_config.SystemParams.REDIS_DB)
	if err != nil {
		redisDb = 0
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", _config.SystemParams.REDIS_HOST, port),
		Password: _config.SystemParams.REDIS_PASSWORD,
		DB:       redisDb,
	})

	var redisClient IRedisClient = &RedisConn{
		Client: rdb,
	}

	return redisClient
}

func (rc *RedisConn) GetData(ctx context.Context, key string) (string, error) {
	return rc.Client.Get(ctx, "key").Result()
}

func (rc *RedisConn) SetData(ctx context.Context, key string, value string, ttl time.Duration) error {
	return rc.Client.Set(ctx, key, value, ttl).Err()
}

func (rc *RedisConn) DelData(ctx context.Context, keyPattern string) error {
	iter := rc.Client.Scan(ctx, 0, keyPattern+"*", 0).Iterator()
	for iter.Next(ctx) {
		err := rc.Client.Del(ctx, iter.Val()).Err()
		if err != nil {
			return err
		}
	}
	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}
