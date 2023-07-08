package uctx

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/dal/models"
)

func User(ctx *gin.Context) models.User {
	if user, ok := ctx.Get("user"); ok {
		return user.(models.User)
	}
	return models.User{}
}

func UserID(ctx *gin.Context) int64 {
	return User(ctx).ID
}

func Name(ctx *gin.Context) string {
	return User(ctx).Name
}

func Phone(ctx *gin.Context) string {
	return User(ctx).Phone
}
