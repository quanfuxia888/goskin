package common

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"quanfuxia/pkg/config"
	"time"
)

type TokenType string

const (
	TokenAccess  TokenType = "access"
	TokenRefresh TokenType = "refresh"
)

type TokenPair struct {
	AccessToken   string `json:"access_token"`
	AccessExpire  int64  `json:"access_expire"` // Unix 时间戳（秒）
	RefreshToken  string `json:"refresh_token"`
	RefreshExpire int64  `json:"refresh_expire"`
}

type CustomClaims struct {
	UserID    int64     `json:"user_id"`
	TokenType TokenType `json:"token_type"`
	JTI       string    `json:"jti"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int64, typ TokenType) (string, *CustomClaims, error) {
	secret := []byte(config.Cfg.JWT.Secret)
	var expire time.Duration

	switch typ {
	case TokenAccess:
		expire = time.Duration(config.Cfg.JWT.AccessExpire) * time.Minute
	case TokenRefresh:
		expire = time.Duration(config.Cfg.JWT.RefreshExpire) * time.Minute
	}

	expireTime := time.Now().Add(expire)
	claims := &CustomClaims{
		UserID:    userID,
		TokenType: typ,
		JTI:       uuid.New().String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secret)
	return tokenStr, claims, err
}

func ParseToken(tokenStr string, expectedType TokenType) (*CustomClaims, error) {
	secret := []byte(config.Cfg.JWT.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims.TokenType != expectedType {
		return nil, errors.New("token type mismatch")
	}
	return claims, nil
}
