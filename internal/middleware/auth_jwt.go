package middleware

import (
	"github.com/gin-gonic/gin"
	"quanfuxia/internal/common"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			common.FailWithStatus(c, common.ErrTokenMissing, 401)
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		claims, err := common.ParseToken(tokenStr, common.TokenAccess)
		if err != nil {
			common.FailWithStatus(c, common.ErrTokenInvalid, 401)
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
