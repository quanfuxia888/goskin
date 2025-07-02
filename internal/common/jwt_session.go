package common

import (
	"context"
	"fmt"
	"time"

	"quanfuxia/pkg/redis"
)

func getRefreshTokenKey(jti string) string {
	return fmt.Sprintf("jwt:refresh:blacklist:%s", jti)
}

func StoreRefreshTokenJTI(jti string, ttl time.Duration) error {
	return redis.Rdb.Set(context.Background(), getRefreshTokenKey(jti), "1", ttl).Err()
}

func IsRefreshTokenRevoked(jti string) bool {
	_, err := redis.Rdb.Get(context.Background(), getRefreshTokenKey(jti)).Result()
	return err == nil
}

func RevokeRefreshToken(jti string) error {
	return redis.Rdb.Set(context.Background(), getRefreshTokenKey(jti), "revoked", time.Hour*24*7).Err()
}
