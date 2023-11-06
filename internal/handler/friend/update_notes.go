package friend

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"github.com/tdfxlyh/go-gin-api/internal/utils/uctx"
)

type UpdateNotesHandler struct {
	Ctx *gin.Context
	Req UpdateNotesReq

	Err error
}

type UpdateNotesReq struct {
	UserID int64  `json:"user_id"`
	Notes  string `json:"notes"`
}

func NewUpdateNotesHandler(ctx *gin.Context) *UpdateNotesHandler {
	return &UpdateNotesHandler{
		Ctx: ctx,
	}

}

func UpdateNotes(ctx *gin.Context) *output.RespStu {
	h := NewUpdateNotesHandler(ctx)

	if msg := h.CheckReqAndUpdate(); msg != "" {
		return output.FailWithMsg(h.Ctx, output.StatusCodeParamsError, msg)
	}
	return output.Success(ctx, nil)
}

func (h *UpdateNotesHandler) CheckReqAndUpdate() string {
	h.Ctx.Bind(&h.Req)
	if h.Req.UserID == 0 || h.Req.Notes == "" {
		return "缺少参数"
	}
	var user models.User
	caller.LyhTestDB.Table(models.TableNameUser).Where("uid=? and status=0", h.Req.UserID).First(&user)
	if user.UID == 0 {
		return "用户不存在"
	}
	var friendRel models.FriendRelation
	caller.LyhTestDB.Table(models.TableNameFriendRelation).Where("user_id=? and other_user_id=? and rela_status=2 and status=0", uctx.UID(h.Ctx), user.UID).First(&friendRel)
	if friendRel.ID == 0 {
		return "请先添加好友"
	}
	caller.LyhTestDB.Table(models.TableNameFriendRelation).Where("id=?", friendRel.ID).Update("notes", h.Req.Notes)
	return ""
}
