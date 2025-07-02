package redis

import (
	"context"
	"errors"
	"time"

	"github.com/bsm/redislock"
)

var locker *redislock.Client

func InitLocker() {
	locker = redislock.New(Rdb)
}

// 获取 Redis 分布式锁（RedLock）
// key: 锁名，ttl: 自动释放时间
func TryLock(ctx context.Context, key string, ttl time.Duration) (*redislock.Lock, error) {
	lock, err := locker.Obtain(ctx, key, ttl, nil)
	if errors.Is(err, redislock.ErrNotObtained) {
		return nil, errors.New("锁已被占用")
	}
	return lock, err
}

// 释放锁（如果业务提前完成）
func ReleaseLock(lock *redislock.Lock) {
	_ = lock.Release(context.Background())
}
