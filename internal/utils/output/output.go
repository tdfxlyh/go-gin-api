package output

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespStu struct{}

func Success(ctx *gin.Context, data interface{}) *RespStu {
	ctx.JSON(http.StatusOK, NewStdResp(data))
	return nil
}

func Fail(ctx *gin.Context, code ErrorCode) *RespStu {
	ctx.JSON(http.StatusOK, NewErrorResp(code, ""))
	return nil
}

func FailWithMsg(ctx *gin.Context, code ErrorCode, customMsg string) *RespStu {
	ctx.JSON(http.StatusOK, NewErrorResp(code, customMsg))
	return nil
}
