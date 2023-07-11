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

func UID(ctx *gin.Context) int64 {
	return User(ctx).UID
}

func UserName(ctx *gin.Context) string {
	return User(ctx).UserName
}

func Phone(ctx *gin.Context) string {
	return User(ctx).Phone
}
