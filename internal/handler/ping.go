package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/internal/model/res"
)

func Ping(ctx *gin.Context) *res.RespStu {
	return res.Success(ctx, gin.H{
		"message": "pong",
	})
}
