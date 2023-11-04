package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/utils"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"github.com/tdfxlyh/go-gin-api/internal/utils/uctx"
	"sync"
	"time"
)

type GetMessageHandler struct {
	Ctx  *gin.Context
	Req  GetMessageReq
	Resp GetMessageResp

	LastTime            time.Time               // 上次的时间
	OtherUserName       string                  // 对方昵称
	MyAvatar            string                  // 我的头像
	OtherAvatar         string                  // 对方头像
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
	Id          int64  `json:"id"`
	IsMe        bool   `json:"is_me"`
	AvatarUrl   string `json:"avatar_url"`
	MessageType int64  `json:"message_type"`
	Content     string `json:"content"`
	StdExtra    string `json:"std_extra"`
	TimeStr     string `json:"time_str"`
	Timestamp   int64  `json:"timestamp"` // 毫秒级别
	WithDraw    int32  `json:"with_draw"` // 是否撤回
}

func NewGetMessageHandler(ctx *gin.Context) *GetMessageHandler {
	h := &GetMessageHandler{
		Ctx:      ctx,
		Req:      GetMessageReq{},
		LastTime: utils.GetYearOf2000(),
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
		return output.FailWithMsg(h.Ctx, output.StatusCodeSeverException, h.Err.Error())
	}
	h.PackData()
	// 成功
	return output.Success(h.Ctx, h.Resp)
}

func (h *GetMessageHandler) CheckReq() {
	h.Ctx.Bind(&h.Req)
	if h.Req.ReceiverUserID == 0 {
		h.Err = fmt.Errorf("receiver_user_id is nil")
	}
	if h.Req.ReceiverUserID == uctx.UID(h.Ctx) {
		h.Err = fmt.Errorf("user_id is repeat")
	}
	switch h.Req.OptType {
	case New100:
		return
	case NewN:
		if h.Req.N == 0 {
			h.Err = fmt.Errorf("n is nil")
			return
		}
	case NewInfo:
		h.LastTime = time.UnixMilli(h.Req.Timestamp)
	case Forward50, NewInfoCount:
		if h.Req.Timestamp == 0 {
			h.Err = fmt.Errorf("timestamp is nil")
			return
		}
	default:
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
	// 查对方昵称
	var friendInfo models.FriendRelation
	caller.LyhTestDB.Table(models.TableNameFriendRelation).Debug().
		Where("user_id=? and other_user_id=? and rela_status=2", uctx.UID(h.Ctx), h.Req.ReceiverUserID).First(&friendInfo)
	if friendInfo.ID == 0 {
		err := fmt.Errorf("friend not found")
		return err
	}
	h.OtherUserName = friendInfo.Notes
	// 查头像
	userList := make([]models.User, 0)
	caller.LyhTestDB.Table(models.TableNameUser).Where("uid=? or uid=?", uctx.UID(h.Ctx), h.Req.ReceiverUserID).Find(&userList)
	if len(userList) < 2 {
		err := fmt.Errorf("user not exist")
		return err
	}
	for _, user := range userList {
		if user.UID == uctx.UID(h.Ctx) {
			h.MyAvatar = utils.GetPic(user.Avatar)
		} else {
			h.OtherAvatar = utils.GetPic(user.Avatar)
			if h.OtherUserName == "" {
				h.OtherUserName = user.UserName
			}
		}
	}
	return nil
}

func (h *GetMessageHandler) GetMsgList() error {
	if h.Req.OptType == New100 {
		h.GetMsgNewN(100)
	} else if h.Req.OptType == NewN {
		h.GetMsgNewN(h.Req.N)
	} else if h.Req.OptType == Forward50 {
		h.GetMsgForward50()
	} else if h.Req.OptType == NewInfo || h.Req.OptType == NewInfoCount {
		h.GetMsgNewInfo()
	}
	return nil
}

func (h *GetMessageHandler) GetMsgNewN(n int64) error {
	caller.LyhTestDB.Table(models.TableNameMessageSingle).
		Where("(sender_user_id=? and receiver_user_id=?) or (sender_user_id=? and receiver_user_id=?)",
			uctx.UID(h.Ctx), h.Req.ReceiverUserID, h.Req.ReceiverUserID, uctx.UID(h.Ctx)).
		Order("create_time desc").Limit(int(n)).
		Find(&h.MsgSingleList)
	return nil
}

func (h *GetMessageHandler) GetMsgForward50() error {
	caller.LyhTestDB.Table(models.TableNameMessageSingle).
		Where("((sender_user_id=? and receiver_user_id=?) or (sender_user_id=? and receiver_user_id=?)) and create_time < ?",
			uctx.UID(h.Ctx), h.Req.ReceiverUserID, h.Req.ReceiverUserID, uctx.UID(h.Ctx), time.UnixMilli(h.Req.Timestamp)).
		Order("create_time desc").Limit(50).
		Find(&h.MsgSingleList)
	return nil
}

func (h *GetMessageHandler) GetMsgNewInfo() error {
	caller.LyhTestDB.Table(models.TableNameMessageSingle).
		Where("((sender_user_id=? and receiver_user_id=?) or (sender_user_id=? and receiver_user_id=?)) and create_time > ?",
			uctx.UID(h.Ctx), h.Req.ReceiverUserID, h.Req.ReceiverUserID, uctx.UID(h.Ctx), time.UnixMilli(h.Req.Timestamp)).
		Order("create_time desc").
		Find(&h.MsgSingleList)
	return nil
}

func (h *GetMessageHandler) PackData() {
	h.Resp.UID = h.Req.ReceiverUserID
	h.Resp.UserName = h.OtherUserName
	h.Resp.Count = int64(len(h.MsgSingleList))
	if h.Req.OptType != NewInfoCount {
		// 先逆序
		h.ReverseArray()
		for _, msg := range h.MsgSingleList {
			// 删除的消息
			if (msg.SenderUserID == uctx.UID(h.Ctx) && msg.SenderStatusInfo == 1) || (msg.ReceiverUserID == uctx.UID(h.Ctx) && msg.ReceiverStatusInfo == 1) {
				continue
			}
			msgItem := &MsgItem{
				Id:          msg.ID,
				IsMe:        false,
				AvatarUrl:   h.OtherAvatar,
				MessageType: msg.MessageType,
				Content:     msg.Content,
				StdExtra:    msg.Extra,
				Timestamp:   msg.CreateTime.UnixMilli(),
				WithDraw:    msg.Withdraw,
			}
			if msg.SenderUserID == uctx.UID(h.Ctx) {
				msgItem.IsMe = true
				msgItem.AvatarUrl = h.MyAvatar
			}
			if utils.IsTimeDifferenceThanNMinute(msg.CreateTime, h.LastTime, 1) {
				if msg.CreateTime.After(utils.GetTodayMidnight()) {
					msgItem.TimeStr = msg.CreateTime.Format("15:04")
				} else if msg.CreateTime.After(utils.GetYesterdayMidnight()) {
					msgItem.TimeStr = msg.CreateTime.Format("昨天 15:04")
				} else if msg.CreateTime.After(utils.GetStartOfYear()) {
					msgItem.TimeStr = msg.CreateTime.Format("1月2日 15:04")
				} else {
					msgItem.TimeStr = msg.CreateTime.Format("2006年1月2日 15:04")
				}
			}
			h.LastTime = msg.CreateTime
			// 如果消息撤回了，不下发内容
			if msg.Withdraw == 1 {
				msgItem.Content = ""
				msgItem.StdExtra = msg.Extra
			}
			h.Resp.MsgList = append(h.Resp.MsgList, msgItem)
		}
	}
	// 最后更新消息为已读
	go h.UpdateMsgRead()
}

// UpdateMsgRead 更新为已读
func (h *GetMessageHandler) UpdateMsgRead() {
	if h.Req.OptType == New100 || h.Req.OptType == NewN || h.Req.OptType == Forward50 || h.Req.OptType == NewInfo {
		for _, msg := range h.MsgSingleList {
			if msg.SenderUserID == h.Req.ReceiverUserID && msg.ReadStatusInfo == 0 {
				h.NeedUpdateMsgIDList = append(h.NeedUpdateMsgIDList, msg.ID)
			}
		}
		caller.LyhTestDB.Debug().Table(models.TableNameMessageSingle).Where("id in (?)", h.NeedUpdateMsgIDList).Update("read_status_info", 1)
	}
}

func (h *GetMessageHandler) ReverseArray() {
	length := len(h.MsgSingleList)
	for i := 0; i < length/2; i++ {
		h.MsgSingleList[i], h.MsgSingleList[length-i-1] = h.MsgSingleList[length-i-1], h.MsgSingleList[i]
	}
}
