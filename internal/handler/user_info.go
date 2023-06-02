package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/cronloader"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"net/http"
)

type UserInfoHandler struct {
	Ctx context.Context

	UserInfoList []*models.User
}

func NewUserInfoHandler(ctx *gin.Context) *UserInfoHandler {
	return &UserInfoHandler{
		Ctx:          ctx,
		UserInfoList: make([]*models.User, 0),
	}
}

func UserInfo(ctx *gin.Context) {
	h := NewUserInfoHandler(ctx)

	h.Process()

	ctx.JSON(http.StatusOK, gin.H{
		"user_list": h.UserInfoList,
	})
}

func (h *UserInfoHandler) Process() {
	h.ReadDataFromDB()
}

func (h *UserInfoHandler) ReadDataFromDB() {
	if cronloader.UserInfoList != nil {
		h.UserInfoList = cronloader.UserInfoList
	}
}
