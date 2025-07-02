package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"quanfuxia/internal/common"
)

func I18nMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := viper.GetString("lang")
		if lang == "" {
			lang = "zh"
		}
		fmt.Println("lang:" + lang)
		if trans, ok := common.Translators[lang]; ok {
			c.Set("trans", trans)
		} else {
			c.Set("trans", common.Translators["zh"])
		}

		c.Next()
	}
}
