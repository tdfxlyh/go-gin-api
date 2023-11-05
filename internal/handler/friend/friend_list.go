package friend

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

type GetFriendListHandler struct {
	Ctx  *gin.Context
	Resp *GetFriendListResp
	Lock sync.RWMutex

	FriendUserInfoMap map[int64]*UserItem
	FriendUserIDs     []int64

	Err error
}

type GetFriendListResp struct {
	UserList []*UserItem `json:"msg_list"`
	Count    int64       `json:"count"`
}

type UserItem struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Pinyin   string `json:"pinyin"`
	Count    int64  `json:"count"`
	CountStr string `json:"count_str"`
}

func NewFriendListHandler(ctx *gin.Context) *GetFriendListHandler {
	return &GetFriendListHandler{
		Ctx:               ctx,
		Lock:              sync.RWMutex{},
		FriendUserIDs:     make([]int64, 0),
		FriendUserInfoMap: make(map[int64]*UserItem),

		Resp: &GetFriendListResp{
			UserList: make([]*UserItem, 0),
		},
	}
}

func GetFriendList(ctx *gin.Context) *output.RespStu {
	h := NewFriendListHandler(ctx)

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

func (h *GetFriendListHandler) GetFriends() {
	friendList := make([]models.FriendRelation, 0)
	h.Err = caller.LyhTestDB.Debug().Table(models.TableNameFriendRelation).Where("status=0 and rela_status=2 and user_id=?", uctx.UID(h.Ctx)).Find(&friendList).Error
	if h.Err != nil {
		fmt.Printf("[GetFriendListHandler-GetFriends] db fail, err=%s\n", h.Err)
		return
	}
	for _, friend := range friendList {
		h.FriendUserIDs = append(h.FriendUserIDs, friend.OtherUserID)
		h.FriendUserInfoMap[friend.OtherUserID] = &UserItem{
			Id:   friend.OtherUserID,
			Name: friend.Notes,
		}
	}
}

func (h *GetFriendListHandler) GetUserInfoAndMsgCount() {
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

func (h *GetFriendListHandler) GetUsersInfo() {
	userList := make([]models.User, 0)
	h.Err = caller.LyhTestDB.Debug().Table(models.TableNameUser).Where("uid in (?)", h.FriendUserIDs).Find(&userList).Error
	if h.Err != nil {
		fmt.Printf("[GetFriendListHandler-GetFriends] db fail, err=%s\n", h.Err)
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
		h.FriendUserInfoMap[user.UID].Pinyin = utils.GetFirstPinYin(h.FriendUserInfoMap[user.UID].Name)
		h.Lock.Unlock()
	}
	return
}

func (h *GetFriendListHandler) GetMsgCount() {
	var wg sync.WaitGroup
	for _, friendID := range h.FriendUserIDs {
		wg.Add(1)
		go func(friendID int64) {
			defer wg.Done()
			msgList := make([]*models.MessageSingle, 0)
			caller.LyhTestDB.Table(models.TableNameMessageSingle).
				Where("sender_user_id=? and receiver_user_id=?",
					friendID, uctx.UID(h.Ctx)).
				Order("create_time desc").Limit(120).
				Find(&msgList)
			for _, msg := range msgList {
				if msg.ReadStatusInfo == 0 {
					h.Lock.Lock()
					h.FriendUserInfoMap[friendID].Count++
					h.Lock.Unlock()
				}
			}
		}(friendID)
	}
	wg.Wait()
}

func (h *GetFriendListHandler) PackData() {
	for _, friendInfo := range h.FriendUserInfoMap {
		if friendInfo.Count > 0 && friendInfo.Count <= 99 {
			friendInfo.CountStr = fmt.Sprintf("%d", friendInfo.Count)
		} else if friendInfo.Count > 99 {
			friendInfo.CountStr = "99+"
		}
		h.Resp.UserList = append(h.Resp.UserList, friendInfo)
	}
	// 按首字符排序
	sort.Slice(h.Resp.UserList, func(i, j int) bool {
		if h.Resp.UserList[i].Pinyin == h.Resp.UserList[j].Pinyin {
			return h.Resp.UserList[i].Id < h.Resp.UserList[j].Id
		}
		return h.Resp.UserList[i].Pinyin < h.Resp.UserList[j].Pinyin
	})
	h.Resp.Count = int64(len(h.Resp.UserList))
}
