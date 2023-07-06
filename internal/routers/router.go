package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/internal/handler"
	"github.com/tdfxlyh/go-gin-api/internal/handler/user"
	"github.com/tdfxlyh/go-gin-api/internal/middleware"
	"github.com/tdfxlyh/go-gin-api/internal/model/res"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/ping", dec(handler.Ping))

	// 注册、登录
	r.POST("/register", dec(user.Register))
	r.POST("/login", dec(user.Login))

	// 获取用户信息
	infoGroup := r.Group("/info", middleware.AuthMiddleware())
	infoGroup.GET("/user_info", dec(user.UserInfo))

	return r
}

// 为了返回体使用return语句，举例: return res.Success(ctx, data)
func dec(f func(ctx *gin.Context) *res.RespStu) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_ = f(ctx) // 原方法
	}
}
