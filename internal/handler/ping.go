package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
