package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/internal/handler"
	"github.com/tdfxlyh/go-gin-api/internal/handler/friend"
	"github.com/tdfxlyh/go-gin-api/internal/handler/message"
	"github.com/tdfxlyh/go-gin-api/internal/handler/user"
	"github.com/tdfxlyh/go-gin-api/internal/middleware"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// 全局捕获异常
	r.Use(middleware.ExceptionMiddleware)

	// ping、注册、登录
	r.GET("/ping", dec(handler.Ping))
	r.POST("/register", dec(user.Register))
	r.POST("/login", dec(user.Login))

	// 获取用户信息
	infoGroup := r.Group("/info", middleware.AuthMiddleware())
	infoGroup.POST("/user_info", dec(user.UserInfo))

	// 消息
	messageGroup := r.Group("/message", middleware.AuthMiddleware())
	messageGroup.POST("/add_message", dec(message.AddMessage))
	messageGroup.POST("/get_message", dec(message.GetMessage))
	messageGroup.POST("/opt_message", dec(message.OptMessage))

	// 朋友信息
	friendGroup := r.Group("/friend", middleware.AuthMiddleware())
	friendGroup.POST("/friend_list", dec(friend.GetFriendList))
	friendGroup.POST("/add_friend", dec(friend.AddFriend))
	friendGroup.POST("/update_notes", dec(friend.UpdateNotes))

	return r
}

// 为了返回体使用return语句，举例: return res.Success(ctx, data)
func dec(f func(ctx *gin.Context) *output.RespStu) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_ = f(ctx) // 原方法
	}
}
