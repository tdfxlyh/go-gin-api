package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/internal/handler"
	"github.com/tdfxlyh/go-gin-api/internal/handler/user"
	"github.com/tdfxlyh/go-gin-api/internal/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/ping", handler.Ping)

	// 注册、登录
	r.POST("/register", user.Register)
	r.POST("/login", user.Login)

	// 获取用户信息
	r.GET("/user_info", middleware.AuthMiddleware(), user.UserInfo)

	return r
}
