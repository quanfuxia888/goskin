package user

import (
	"github.com/gin-gonic/gin"
	"quanfuxia/internal/common"
	"quanfuxia/internal/service"
)

type UserHandler struct {
	Svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{Svc: svc}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.HandleValidationError(c, err)
		return
	}

	err := h.Svc.Register(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		common.Fail(c, common.ErrInternal)
		return
	}
	common.Success(c, gin.H{"msg": "注册成功"})
}

func (h *UserHandler) UserInfo(c *gin.Context) {
	userID := c.GetInt64("userID")
	c.JSON(200, gin.H{"code": 0, "user_id": userID})
}

func (h *UserHandler) Refresh(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		common.HandleValidationError(c, err)
		return
	}

	claims, err := common.ParseToken(req.RefreshToken, common.TokenRefresh)
	if err != nil || common.IsRefreshTokenRevoked(claims.JTI) {
		common.Fail(c, common.ErrTokenInvalid)
		return
	}

	_ = common.RevokeRefreshToken(claims.JTI)

	accessToken, _, _ := common.GenerateToken(claims.UserID, common.TokenAccess)
	refreshToken, refreshClaims, _ := common.GenerateToken(claims.UserID, common.TokenRefresh)

	common.SuccessTokenPair(c, accessToken, refreshToken, refreshClaims.JTI)
}

// ===================== handler/user/handler.go 中 Login 接口 =====================

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		common.HandleValidationError(c, err)
		return
	}

	userID := int64(1001)
	accessToken, _, _ := common.GenerateToken(userID, common.TokenAccess)
	refreshToken, refreshClaims, _ := common.GenerateToken(userID, common.TokenRefresh)

	common.SuccessTokenPair(c, accessToken, refreshToken, refreshClaims.JTI)
}
