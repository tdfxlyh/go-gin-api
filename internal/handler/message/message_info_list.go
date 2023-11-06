package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/utils"
	"github.com/tdfxlyh/go-gin-api/internal/utils/output"
	"github.com/tdfxlyh/go-gin-api/internal/utils/uctx"
	"sort"
	"sync"
)

type GetMessageInfoListHandler struct {
	Ctx  *gin.Context
	Resp *GetMessageInfoListResp
	Lock sync.RWMutex

	FriendUserInfoMap map[int64]*UserItem
	FriendUserIDs     []int64

	Err error
}

type GetMessageInfoListResp struct {
	UserList []*UserItem `json:"msg_list"`
	Count    int64       `json:"count"`
}

type UserItem struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Count     int64  `json:"count"`
	CountStr  string `json:"count_str"`
	Timestamp int64  `json:"timestamp"` // 毫秒级别
	TimeStr   string `json:"time_str"`
	Desc      string `json:"desc"` // 最新一句话
}

func NewGetMessageInfoListHandler(ctx *gin.Context) *GetMessageInfoListHandler {
	return &GetMessageInfoListHandler{
		Ctx:               ctx,
		Lock:              sync.RWMutex{},
		FriendUserIDs:     make([]int64, 0),
		FriendUserInfoMap: make(map[int64]*UserItem),

		Resp: &GetMessageInfoListResp{
			UserList: make([]*UserItem, 0),
		},
	}
}

func GetMessageInfoList(ctx *gin.Context) *output.RespStu {
	h := NewGetMessageInfoListHandler(ctx)

	// 获取好友信息
	if h.GetFriends(); h.Err != nil {
		return output.Fail(ctx, output.StatusCodeDBError)
	}

	// 获取好友用户信息 和消息未读数
	if h.GetUserInfoAndMsgCount(); h.Err != nil {
		return output.Fail(ctx, output.StatusCodeDBError)
	}

	// 打包数据
	h.PackData()

	return output.Success(ctx, h.Resp)
}

func (h *GetMessageInfoListHandler) GetFriends() {
	GetMessageInfoList := make([]models.FriendRelation, 0)
	h.Err = caller.LyhTestDB.Debug().Table(models.TableNameFriendRelation).Where("rela_status=2 and user_id=? and status=0", uctx.UID(h.Ctx)).Find(&GetMessageInfoList).Error
	if h.Err != nil {
		fmt.Printf("[GetMessageInfoListHandler-GetFriends] db fail, err=%s\n", h.Err)
		return
	}
	for _, friend := range GetMessageInfoList {
		h.FriendUserIDs = append(h.FriendUserIDs, friend.OtherUserID)
		h.FriendUserInfoMap[friend.OtherUserID] = &UserItem{
			Id:   friend.OtherUserID,
			Name: friend.Notes,
		}
	}
}

func (h *GetMessageInfoListHandler) GetUserInfoAndMsgCount() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 获取好友用户信息
		h.GetUsersInfo()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 获取消息未读数
		h.GetMsgCount()
	}()
	wg.Wait()
}

func (h *GetMessageInfoListHandler) GetUsersInfo() {
	userList := make([]models.User, 0)
	h.Err = caller.LyhTestDB.Debug().Table(models.TableNameUser).Where("uid in (?) and status=0", h.FriendUserIDs).Find(&userList).Error
	if h.Err != nil {
		fmt.Printf("[GetMessageInfoListHandler-GetFriends] db fail, err=%s\n", h.Err)
		return
	}
	for _, user := range userList {
		if h.FriendUserInfoMap[user.UID] == nil {
			continue
		}
		if h.FriendUserInfoMap[user.UID].Name == "" {
			h.Lock.Lock()
			h.FriendUserInfoMap[user.UID].Name = user.UserName
			h.Lock.Unlock()
		}
		h.Lock.Lock()
		h.FriendUserInfoMap[user.UID].Avatar = utils.GetPic(user.Avatar)
		h.Lock.Unlock()
	}
}

func (h *GetMessageInfoListHandler) GetMsgCount() {
	var wg sync.WaitGroup
	for _, friendID := range h.FriendUserIDs {
		wg.Add(1)
		go func(friendID int64) {
			defer wg.Done()
			msgList := make([]*models.MessageSingle, 0)
			caller.LyhTestDB.Table(models.TableNameMessageSingle).
				Where("((sender_user_id=? and receiver_user_id=?) or (sender_user_id=? and receiver_user_id=?)) and status=0",
					friendID, uctx.UID(h.Ctx), uctx.UID(h.Ctx), friendID).
				Order("create_time desc").Limit(150).
				Find(&msgList)
			for _, msg := range msgList {
				// 删除的消息
				if (msg.SenderUserID == uctx.UID(h.Ctx) && msg.SenderStatusInfo == 1) || (msg.ReceiverUserID == uctx.UID(h.Ctx) && msg.ReceiverStatusInfo == 1) {
					continue
				}
				if h.FriendUserInfoMap[friendID].Timestamp == 0 {
					timeStr := ""
					if msg.CreateTime.After(utils.GetTodayMidnight()) {
						timeStr = msg.CreateTime.Format("15:04")
					} else if msg.CreateTime.After(utils.GetYesterdayMidnight()) {
						timeStr = msg.CreateTime.Format("昨天")
					} else if msg.CreateTime.After(utils.GetStartOfYear()) {
						timeStr = msg.CreateTime.Format("1月2日")
					} else {
						timeStr = msg.CreateTime.Format("2006年1月2日")
					}
					h.Lock.Lock()
					h.FriendUserInfoMap[friendID].Timestamp = msg.CreateTime.UnixMilli()
					h.FriendUserInfoMap[friendID].Desc = msg.Content
					if msg.Withdraw == 1 {
						h.FriendUserInfoMap[friendID].Desc = ">>对方撤回了一条消息"
						if msg.SenderUserID == uctx.UID(h.Ctx) {
							h.FriendUserInfoMap[friendID].Desc = ">>我撤回了一条消息"
						}
					}
					h.FriendUserInfoMap[friendID].TimeStr = timeStr
					h.Lock.Unlock()
				}
				if msg.ReadStatusInfo == 0 && msg.SenderUserID == friendID {
					h.Lock.Lock()
					h.FriendUserInfoMap[friendID].Count++
					h.Lock.Unlock()
				}
			}
		}(friendID)
	}
	wg.Wait()
}

func (h *GetMessageInfoListHandler) PackData() {
	for _, friendInfo := range h.FriendUserInfoMap {
		if friendInfo.Count > 0 && friendInfo.Count <= 99 {
			friendInfo.CountStr = fmt.Sprintf("%d", friendInfo.Count)
		} else if friendInfo.Count > 99 {
			friendInfo.CountStr = "99+"
		}
		h.Resp.UserList = append(h.Resp.UserList, friendInfo)
	}
	// 按时间戳排序
	sort.Slice(h.Resp.UserList, func(i, j int) bool {
		if h.Resp.UserList[i].Timestamp == h.Resp.UserList[j].Timestamp {
			return h.Resp.UserList[i].Id < h.Resp.UserList[j].Id
		}
		return h.Resp.UserList[i].Timestamp > h.Resp.UserList[j].Timestamp
	})
	h.Resp.Count = int64(len(h.Resp.UserList))
}
