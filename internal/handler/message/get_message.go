package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/utils"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"sync"
)

type GetMessageHandler struct {
	Ctx  *gin.Context
	Req  GetMessageReq
	Resp GetMessageResp

	MsgSingleList       []*models.MessageSingle // 数据库读出的记录
	NeedUpdateMsgIDList []int64                 // 需要更新为已读的消息id列表

	Err error
}

const (
	New100       = 1 // 最新100条
	NewN         = 2 // 最新n条消息
	Forward50    = 3 // 往前50条消息
	NewInfo      = 4 // 新消息
	NewInfoCount = 5 // 新消息个数
)

type GetMessageReq struct {
	OptType        int64 `json:"opt_type"`
	ReceiverUserID int64 `json:"receiver_user_id"`
	Timestamp      int64 `json:"timestamp"`
	N              int64 `json:"n"`
}

type GetMessageResp struct {
	UID      int64      `json:"uid"`
	UserName string     `json:"user_name"`
	MsgList  []*MsgItem `json:"msg_list"`
	Count    int64      `json:"count"`
}
type MsgItem struct {
	IsMe        bool   `json:"is_me"`
	AvatarUrl   string `json:"avatar_url"`
	MessageType int64  `json:"message_type"`
	Content     string `json:"content"`
	StdExtra    string `json:"std_extra"`
	TimeStr     string `json:"time_str"`
	Timestamp   int64  `json:"timestamp"` // 毫秒级别
	WithDraw    int64  `json:"with_draw"` // 是否撤回
}

func NewGetMessageHandler(ctx *gin.Context) *GetMessageHandler {
	h := &GetMessageHandler{
		Ctx: ctx,
		Req: GetMessageReq{},
	}
	h.Resp = GetMessageResp{}
	h.Resp.MsgList = make([]*MsgItem, 0)
	return h
}

func GetMessage(ctx *gin.Context) *output.RespStu {
	h := NewGetMessageHandler(ctx)
	// 校验参数
	if h.CheckReq(); h.Err != nil {
		fmt.Printf("[GetMessage-CheckReq] err=%s\n", utils.GetStuStr(h.Err))
		return output.FailWithMsg(h.Ctx, output.StatusCodeParamsError, h.Err.Error())
	}
	// 加载数据
	if h.LoadDataFromDB(); h.Err != nil {
		fmt.Printf("[GetMessage-Process] err=%s\n", utils.GetStuStr(h.Err))
		return output.Fail(h.Ctx, output.StatusCodeSeverException)
	}
	h.PackData()
	// 成功
	return output.Success(h.Ctx, h.Resp)
}

func (h *GetMessageHandler) CheckReq() {
	h.Ctx.Bind(&h.Req)
	if h.Req.OptType == New100 {
		if h.Req.ReceiverUserID == 0 {
			h.Err = fmt.Errorf("receiver_user_id is nil")
			return
		}
	} else if h.Req.OptType == NewN {
		if h.Req.ReceiverUserID == 0 || h.Req.N == 0 {
			h.Err = fmt.Errorf("receiver_user_id or n is nil")
			return
		}
	} else if h.Req.OptType == Forward50 || h.Req.OptType == NewInfo || h.Req.OptType == NewInfoCount {
		if h.Req.ReceiverUserID == 0 || h.Req.Timestamp == 0 {
			h.Err = fmt.Errorf("receiver_user_id or n is nil")
			return
		}
	} else {
		h.Err = fmt.Errorf("opt_type value is error")
		return
	}
}

func (h *GetMessageHandler) LoadDataFromDB() {
	errs := make([]error, 0)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 拿到用户信息
		if err := h.GetUserInfo(); err != nil {
			errs = append(errs, err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 拿到消息
		if err := h.GetMsgList(); err != nil {
			errs = append(errs, err)
		}
	}()
	wg.Wait()
	// 处理消息
	if len(errs) > 0 {
		h.Err = errs[0]
	}
}

func (h *GetMessageHandler) GetUserInfo() error {
	// 根据user_id查用户信息
	// 。。。
	return nil
}

func (h *GetMessageHandler) GetMsgList() error {
	if h.Req.OptType == New100 {
		h.GetMsgNewN(100)
	} else if h.Req.OptType == NewN {
		h.GetMsgNewN(h.Req.N)
	} else if h.Req.OptType == Forward50 {
		h.GetMsgForward50()
	} else if h.Req.OptType == NewInfo {
		h.GetMsgNewInfo()
	} else if h.Req.OptType == NewInfoCount {

	}
	return nil
}

func (h *GetMessageHandler) GetMsgNewN(n int64) error {
	// 获取消息，并更新为已读
	return nil
}

func (h *GetMessageHandler) GetMsgForward50() error {
	// 获取消息，并更新为已读
	return nil
}

func (h *GetMessageHandler) GetMsgNewInfo() error {
	// 获取消息，并更新为已读
	return nil
}

func (h *GetMessageHandler) GetMsgNewInfoCount() error {
	// 获取count
	return nil
}

func (h *GetMessageHandler) PackData() {

	// 最后更新消息为已读
	go h.UpdateMsgRead()
}

// UpdateMsgRead 更新为已读
func (h *GetMessageHandler) UpdateMsgRead() {
	caller.LyhTestDB.Debug().Table(models.TableNameMessageSingle).Where("id in (?)", h.NeedUpdateMsgIDList).Update("read_status_info", 1)
}
