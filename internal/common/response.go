package common

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"quanfuxia/pkg/config"
	"time"
)

// Response 统一响应格式
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 成功响应
func Success(c *gin.Context, data interface{}) {
	//lang := c.GetHeader("Accept-Language") // 获取请求的语言
	lang := viper.GetString("lang") // 获取请求的语言
	// 使用 universal-translator 获取语言翻译
	msg := Translate(lang, "ok") // "ok" 是我们错误码中注册的 Key
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess.Code,
		Msg:  msg,
		Data: data,
	})
}

// 错误响应
func Fail(c *gin.Context, err ErrCode) {
	//lang := c.GetHeader("Accept-Language") // 获取请求的语言
	lang := viper.GetString("lang") // 获取请求的语言
	// 动态翻译错误信息
	msg := Translate(lang, err.MsgKey) // 使用翻译器转换 MsgKey
	c.JSON(http.StatusOK, Response{
		Code: err.Code,
		Msg:  msg,
	})
}

// 错误响应（带自定义 HTTP 状态码）
func FailWithStatus(c *gin.Context, err ErrCode, status int) {
	lang := viper.GetString("lang") // 获取请求的语言
	// 动态翻译错误信息
	msg := Translate(lang, err.MsgKey)
	c.JSON(status, Response{
		Code: err.Code,
		Msg:  msg,
	})
}

func SuccessTokenPair(c *gin.Context, accessToken string, refreshToken string, refreshJTI string) {
	now := time.Now()
	accessExpire := now.Add(time.Minute * time.Duration(config.Cfg.JWT.AccessExpire)).Unix()
	refreshExpire := now.Add(time.Minute * time.Duration(config.Cfg.JWT.RefreshExpire)).Unix()

	_ = StoreRefreshTokenJTI(refreshJTI, time.Duration(config.Cfg.JWT.RefreshExpire)*time.Minute)
	Success(c, gin.H{
		"access_token":   accessToken,
		"access_expire":  accessExpire,
		"refresh_token":  refreshToken,
		"refresh_expire": refreshExpire,
	})
}
