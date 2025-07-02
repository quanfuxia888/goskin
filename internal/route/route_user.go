package route

import (
	"github.com/gin-gonic/gin"
	"quanfuxia/internal/api/user"
	"quanfuxia/internal/middleware"
	"quanfuxia/internal/repository"
	"quanfuxia/internal/service"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userRepo := repository.NewUserRepo()
	userSvc := service.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userSvc)

	group := rg.Group("/user")
	group.POST("/register", userHandler.Register)
	group.POST("/login", userHandler.Login)

	authGroup := group.Group("")
	authGroup.Use(middleware.JWTAuthMiddleware())
	authGroup.GET("/info", userHandler.UserInfo)
	authGroup.GET("/refresh", userHandler.Refresh)

}
