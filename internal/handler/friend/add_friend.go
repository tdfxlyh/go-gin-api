package friend

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/model/dto/dto_friend"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"github.com/tdfxlyh/go-gin-api/internal/utils/uctx"
)

type AddFriendHandler struct {
	Ctx *gin.Context
	Req dto_friend.AddFriendReq

	Err error
}

func NewAddFriendHandler(ctx *gin.Context) *AddFriendHandler {
	return &AddFriendHandler{
		Ctx: ctx,
	}

}

func AddFriend(ctx *gin.Context) *output.RespStu {
	h := NewAddFriendHandler(ctx)

	if msg := h.CheckReqAndAdd(); msg != "" {
		return output.FailWithMsg(h.Ctx, output.StatusCodeParamsError, msg)
	}

	return output.Success(ctx, nil)
}

func (h *AddFriendHandler) CheckReqAndAdd() string {
	h.Ctx.Bind(&h.Req)
	if h.Req.Phone == "" {
		return "缺少手机号"
	}
	var user models.User
	caller.LyhTestDB.Table(models.TableNameUser).Where("phone=? and status=0", h.Req.Phone).First(&user)
	if user.UID == 0 {
		return "用户不存在"
	}
	if user.UID == uctx.UID(h.Ctx) {
		return "不能添加自己为好友"
	}
	friendList := make([]models.FriendRelation, 0)
	caller.LyhTestDB.Table(models.TableNameFriendRelation).Where("((user_id=? and other_user_id=? and rela_status=2) or (user_id=? and other_user_id=? and rela_status=2)) and status=0", uctx.UID(h.Ctx), user.UID, user.UID, uctx.UID(h.Ctx)).Find(&friendList)
	if len(friendList) > 0 {
		return "已经是好友了"
	}
	rel1 := &models.FriendRelation{
		UserID:      uctx.UID(h.Ctx),
		OtherUserID: user.UID,
		RelaStatus:  2,
	}
	rel2 := &models.FriendRelation{
		UserID:      user.UID,
		OtherUserID: uctx.UID(h.Ctx),
		RelaStatus:  2,
	}
	caller.LyhTestDB.Create(rel1)
	caller.LyhTestDB.Create(rel2)
	return ""
}
