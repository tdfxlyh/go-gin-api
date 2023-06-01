package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/internal/handler"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/ping", handler.Ping)
	r.GET("/user_info", handler.UserInfo)

	return r
}
