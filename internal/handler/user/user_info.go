package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/cronloader"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
)

type UserInfoHandler struct {
	Ctx *gin.Context

	UserList map[string][]*models.User
}

func NewUserInfoHandler(ctx *gin.Context) *UserInfoHandler {
	return &UserInfoHandler{
		Ctx:      ctx,
		UserList: make(map[string][]*models.User),
	}
}

func UserInfo(ctx *gin.Context) *output.RespStu {
	h := NewUserInfoHandler(ctx)

	h.GetData()

	return output.Success(ctx, h.UserList)
}

func (h *UserInfoHandler) GetData() {
	if cronloader.UserInfoList != nil {
		h.UserList = map[string][]*models.User{
			"tab_list": cronloader.UserInfoList,
		}
	}
}
