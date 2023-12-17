package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/model/dto/dto_user"
	"github.com/tdfxlyh/go-gin-api/internal/utils"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"github.com/tdfxlyh/go-gin-api/internal/utils/uctx"
	"golang.org/x/crypto/bcrypt"
)

type OptUserHandler struct {
	Ctx *gin.Context
	Req dto_user.OptUserReq

	Err error
}

const (
	OptAvatar   = 1
	OptUsername = 2
	OptPassword = 3
)

func NewOptUserHandler(ctx *gin.Context) *OptUserHandler {
	return &OptUserHandler{
		Ctx: ctx,
	}
}

func OptUser(ctx *gin.Context) *output.RespStu {
	h := NewOptUserHandler(ctx)
	// 校验参数
	if msg := h.CheckReq(); msg != "" {
		return output.FailWithMsg(h.Ctx, output.StatusCodeParamsError, msg)
	}
	// 更新数据
	if h.UpdateUserInfo(); h.Err != nil {
		caller.Logger.Warn(fmt.Sprintf("[OptUser-UpdateUserInfo] err=%s\n", h.Err))
		return output.Fail(h.Ctx, output.StatusCodeDBError)
	}
	return output.Success(h.Ctx, nil)
}

func (h *OptUserHandler) CheckReq() string {
	h.Ctx.Bind(&h.Req)
	switch h.Req.OptType {
	case OptAvatar:
	case OptUsername:
	case OptPassword:
	default:
		return "miss opt_type"
	}
	if h.Req.OptType == OptAvatar {
		if h.Req.Avatar == "" {
			return "miss avatar"
		}
		if !utils.CheckOnlinePic(h.Req.Avatar) {
			return "图片链接不合法"
		}
	}
	if h.Req.OptType == OptUsername && h.Req.UserName == "" {
		return "miss user_name"
	}
	if h.Req.OptType == OptPassword && h.Req.Password == "" {
		return "miss password"
	}
	return ""
}

func (h *OptUserHandler) UpdateUserInfo() {
	if h.Req.OptType == OptAvatar {
		h.Err = caller.LyhTestDB.Table(models.TableNameUser).Where("uid=?", uctx.UID(h.Ctx)).Update("avatar", h.Req.Avatar).Error
	}
	if h.Req.OptType == OptUsername {
		h.Err = caller.LyhTestDB.Table(models.TableNameUser).Where("uid=?", uctx.UID(h.Ctx)).Update("user_name", h.Req.UserName).Error
	}
	if h.Req.OptType == OptPassword {
		hasePassword, err := bcrypt.GenerateFromPassword([]byte(h.Req.Password), bcrypt.DefaultCost)
		if err != nil {
			h.Err = err
			return
		}
		h.Err = caller.LyhTestDB.Table(models.TableNameUser).Where("uid=?", uctx.UID(h.Ctx)).Update("password", string(hasePassword)).Error
	}
}
