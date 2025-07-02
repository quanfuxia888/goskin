package common

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// 用于绑定 gin 的验证错误处理，自动翻译首条错误信息
func HandleValidationError(c *gin.Context, err error) {
	// 从 gin 上下文中获取翻译器（必须提前在中间件中注入）
	trans, ok := c.Get("trans")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Invalid request",
		})
		return
	}
	translator := trans.(ut.Translator)
	// 解析 validator 错误
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求参数错误",
		})
		return
	}

	// 取第一条翻译后的错误信息返回
	msg := errs[0].Translate(translator)

	c.JSON(http.StatusBadRequest, gin.H{
		"code": 400,
		"msg":  msg,
	})
}
