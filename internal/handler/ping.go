package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
)

func Ping(ctx *gin.Context) *output.RespStu {
	return output.Success(ctx, gin.H{
		"message": "pong",
	})
}
