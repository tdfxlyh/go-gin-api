package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/utils"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"github.com/tdfxlyh/go-gin-api/internal/utils/uctx"
	"io/fs"
	"strings"
	"time"
)

type AddMessageHandler struct {
	Ctx  *gin.Context
	Req  AddMessageReq
	Resp interface{}

	Err error
}

type AddMessageReq struct {
	MessageType    int64   `json:"message_type"`
	Content        string  `json:"content"`
	ReceiverUserID int64   `json:"receiver_user_id"`
	Timestamp      int64   `json:"timestamp"`
	File           fs.File `json:"file"`
}

const (
	ContentText  = 1 // 文本
	ContentPic   = 2 // 图片
	ContentAudio = 3 // 音频
	ContentVideo = 4 // 视频
	ContentFile  = 5 // 文件
)

func NewAddMessageHandler(ctx *gin.Context) *AddMessageHandler {
	return &AddMessageHandler{
		Ctx: ctx,
	}
}

func AddMessage(ctx *gin.Context) *output.RespStu {
	h := NewAddMessageHandler(ctx)
	// 校验参数
	if h.CheckReq(); h.Err != nil {
		fmt.Printf("[AddMessage-CheckReq] err=%s\n", utils.GetStuStr(h.Err))
		return output.FailWithMsg(h.Ctx, output.StatusCodeParamsError, h.Err.Error())
	}

	// 写入数据
	if h.AddMessageToDB(); h.Err != nil {
		fmt.Printf("[AddMessage-AddMessageToDB] err=%s\n", utils.GetStuStr(h.Err))
		return output.FailWithMsg(h.Ctx, output.StatusCodeDBError, h.Err.Error())
	}
	// 查最新消息
	if h.GetNewInfo(); h.Err != nil {
		fmt.Printf("[AddMessage-GetNewInfo] err=%s\n", utils.GetStuStr(h.Err))
		return output.FailWithMsg(h.Ctx, output.StatusCodeDBError, h.Err.Error())
	}

	return output.Success(h.Ctx, h.Resp)
}

func (h *AddMessageHandler) CheckReq() {
	h.Ctx.Bind(&h.Req)
	if h.Req.ReceiverUserID == 0 {
		h.Err = fmt.Errorf("receiver_user_id is nil")
	}
	if h.Req.ReceiverUserID == uctx.UID(h.Ctx) {
		h.Err = fmt.Errorf("user_id is repeat")
	}
	if h.Req.Content == "" || strings.TrimSpace(h.Req.Content) == "" {
		h.Err = fmt.Errorf("content is nil")
	}
	switch h.Req.MessageType {
	case ContentText, ContentPic, ContentAudio, ContentVideo, ContentFile:
		return
	default:
		h.Err = fmt.Errorf("message_type value is error")
		return
	}
}

func (h *AddMessageHandler) AddMessageToDB() {
	var message *models.MessageSingle
	if h.Req.MessageType == ContentText {
		message = &models.MessageSingle{
			SenderUserID:   uctx.UID(h.Ctx),
			ReceiverUserID: h.Req.ReceiverUserID,
			MessageType:    h.Req.MessageType,
			Content:        h.Req.Content,
		}
	}
	if message == nil {
		return
	}
	caller.LyhTestDB.Table(models.TableNameMessageSingle).Create(message)
}

func (h *AddMessageHandler) GetNewInfo() {
	getMessageHandler := NewGetMessageHandler(h.Ctx)
	optType := NewN
	if h.Req.Timestamp != 0 { // 没有时间戳返回最新10条记录
		optType = NewInfo
		getMessageHandler.LastTime = time.UnixMilli(h.Req.Timestamp)
	}
	getMessageHandler.Req = GetMessageReq{
		OptType:        int64(optType),
		ReceiverUserID: h.Req.ReceiverUserID,
		Timestamp:      h.Req.Timestamp,
		N:              10,
	}
	// 加载数据
	if getMessageHandler.LoadDataFromDB(); h.Err != nil {
		fmt.Printf("[AddMessageHandler-GetMessage-Process] err=%s\n", h.Err)
	}
	getMessageHandler.PackData()
	h.Resp = getMessageHandler.Resp
}
