package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/internal/utils"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"net/http"
)

// ExceptionMiddleware 异常中间件: 捕获全局异常
func ExceptionMiddleware(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 简单返回友好提示，具体可自定义发生错误后处理逻辑
			// ...
			ctx.JSON(http.StatusInternalServerError, output.NewErrorResp(output.StatusCodeSeverException, utils.GetStuStr(err)))
			ctx.Abort()
		}
	}()
	ctx.Next()
}
