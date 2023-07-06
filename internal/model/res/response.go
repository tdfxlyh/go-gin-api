package res

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"net/http"
)

type RespStu struct{}

func Success(ctx *gin.Context, data interface{}) *RespStu {
	ctx.JSON(http.StatusOK, output.NewStdResp(data))
	return nil
}

func Fail(ctx *gin.Context, code output.ErrorCode) *RespStu {
	ctx.JSON(http.StatusOK, output.NewErrorResp(code, ""))
	return nil
}

func FailWithMsg(ctx *gin.Context, code output.ErrorCode, customMsg string) *RespStu {
	ctx.JSON(http.StatusOK, output.NewErrorResp(code, customMsg))
	return nil
}
