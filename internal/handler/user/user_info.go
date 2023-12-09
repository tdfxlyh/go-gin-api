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
)

const (
	UserInfoSelf  = 1 // 自己的信息
	UserInfoOther = 2 // 其他人的信息
)

type UserInfoHandler struct {
	Ctx  *gin.Context
	Req  dto_user.UserInfoReq
	Resp *dto_user.UserInfoResp

	Err error
}

func NewUserInfoHandler(ctx *gin.Context) *UserInfoHandler {
	return &UserInfoHandler{
		Ctx: ctx,

		Resp: &dto_user.UserInfoResp{},
	}
}

func UserInfo(ctx *gin.Context) *output.RespStu {
	h := NewUserInfoHandler(ctx)

	if h.CheckReq(); h.Err != nil {
		fmt.Println(fmt.Sprintf("[UserInfo-CheckReq] fail, err=%s", h.Err))
		return output.Fail(h.Ctx, output.StatusCodeParamsError)
	}
	if h.GetData(); h.Err != nil {
		fmt.Println(fmt.Sprintf("[UserInfo-GetData] fail, err=%s", h.Err))
		return output.Fail(h.Ctx, output.StatusCodeDBError)
	}
	return output.Success(ctx, h.Resp)
}

func (h *UserInfoHandler) CheckReq() {
	h.Ctx.Bind(&h.Req)
	if h.Req.OptType != UserInfoSelf && h.Req.OptType != UserInfoOther {
		h.Err = fmt.Errorf("opt_type is err")
		return
	}
	if h.Req.OptType == UserInfoOther && h.Req.UID == 0 {
		h.Err = fmt.Errorf("miss uid")
		return
	}
}

func (h *UserInfoHandler) GetData() {
	uid := uctx.UID(h.Ctx)
	if h.Req.OptType == UserInfoOther {
		uid = h.Req.UID
	}
	userInfo := models.User{}
	caller.LyhTestDB.Table(models.TableNameUser).Debug().Where("uid=? and status=0", uid).First(&userInfo)
	if userInfo.UID == 0 {
		h.Err = fmt.Errorf("user not found")
		return
	}
	h.Resp = &dto_user.UserInfoResp{
		UID:      userInfo.UID,
		UserName: userInfo.UserName,
		Avatar:   utils.GetPic(userInfo.Avatar),
	}
}
