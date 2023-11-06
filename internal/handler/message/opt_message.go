package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"github.com/tdfxlyh/go-gin-api/internal/utils/uctx"
)

type OptMessageHandler struct {
	Ctx *gin.Context
	Req OptMessageReq
	Msg models.MessageSingle

	Err error
}

const (
	MsgRead     = 1 // 更改已读
	MsgDel      = 2 // 删除消息
	MsgWithdraw = 3 // 撤回消息
)

type OptMessageReq struct {
	OptType   int64 `json:"opt_type"`
	MessageID int64 `json:"message_id"`
}

func NewOptMessageHandler(ctx *gin.Context) *OptMessageHandler {
	return &OptMessageHandler{
		Ctx: ctx,
	}
}

func OptMessage(ctx *gin.Context) *output.RespStu {
	h := NewOptMessageHandler(ctx)
	// 参数校验
	if msg := h.CheckReq(); msg != "" {
		return output.FailWithMsg(h.Ctx, output.StatusCodeParamsError, msg)
	}
	// 操作消息
	if h.OptMsg(); h.Err != nil {
		fmt.Printf("[OptMessage] db fail, err=%s\n", h.Err)
		return output.Fail(h.Ctx, output.StatusCodeDBError)
	}
	return output.Success(h.Ctx, nil)
}

func (h *OptMessageHandler) CheckReq() string {
	h.Ctx.Bind(&h.Req)
	if h.Req.OptType != MsgRead && h.Req.OptType != MsgDel && h.Req.OptType != MsgWithdraw {
		return "操作类型错误"
	}
	if h.Req.MessageID == 0 {
		return "miss message_id"
	}
	caller.LyhTestDB.Debug().Table(models.TableNameMessageSingle).Where("id=? and status=0", h.Req.MessageID).First(&h.Msg)
	if h.Msg.ID == 0 {
		return "msg not found"
	}
	if h.Req.OptType == MsgWithdraw && h.Msg.SenderUserID != uctx.User(h.Ctx).UID {
		return "not operate other user msg"
	}
	return ""
}

func (h *OptMessageHandler) OptMsg() {
	switch h.Req.OptType {
	case MsgRead:
		h.OptMsgRead()
	case MsgDel:
		h.OptMsgDel()
	case MsgWithdraw:
		h.OptMsgWithdraw()
	}
}

// OptMsgRead 消息标记为已读
func (h *OptMessageHandler) OptMsgRead() {
	h.Err = caller.LyhTestDB.Debug().Table(models.TableNameMessageSingle).Where("id=? and receiver_user_id=?", h.Req.MessageID, uctx.User(h.Ctx).UID).Update("read_status_info", 1).Error
}

// OptMsgDel 删除消息
func (h *OptMessageHandler) OptMsgDel() {
	if h.Msg.SenderUserID == uctx.User(h.Ctx).UID { // 如果我是消息发送者
		h.Err = caller.LyhTestDB.Debug().Table(models.TableNameMessageSingle).Where("id=?", h.Req.MessageID).Update("sender_status_info", 1).Error
	} else if h.Msg.ReceiverUserID == uctx.User(h.Ctx).UID { // 如果我是消息接收者
		h.Err = caller.LyhTestDB.Debug().Table(models.TableNameMessageSingle).Where("id=?", h.Req.MessageID).Update("receiver_status_info", 1).Error
	}
}

// OptMsgWithdraw 撤回消息
func (h *OptMessageHandler) OptMsgWithdraw() {
	h.Err = caller.LyhTestDB.Debug().Table(models.TableNameMessageSingle).Where("id=?", h.Req.MessageID).Update("withdraw", 1).Error
}
