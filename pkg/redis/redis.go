package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"quanfuxia/pkg/config"
)

var Rdb *redis.Client

func Init() {
	cfg := config.Cfg.Redis

	Rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// 测试连接
	if err := Rdb.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Sprintf("Redis连接失败: %v", err))
	}

	fmt.Println("✅ Redis 已连接")
}
