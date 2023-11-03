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
)

type GetFriendListHandler struct {
	Ctx  *gin.Context
	Resp *GetFriendListResp

	FriendUserInfoMap map[int64]*UserItem
	FriendUserIDs     []int64

	Err error
}

type GetFriendListResp struct {
	UserList []*UserItem `json:"msg_list"`
	Count    int64       `json:"count"`
}

type UserItem struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Pinyin string `json:"pinyin"`
}

func NewFriendListHandler(ctx *gin.Context) *GetFriendListHandler {
	return &GetFriendListHandler{
		Ctx:               ctx,
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
	// 获取好友用户信息
	if h.GetUsersInfo(); h.Err != nil {
		return output.Fail(ctx, output.StatusCodeDBError)
	}
	// 打包数据
	h.PackData()

	return output.Success(ctx, h.Resp)
}

func (h *GetFriendListHandler) GetFriends() {
	friendList := make([]models.FriendRelation, 0)
	h.Err = caller.LyhTestDB.Debug().Table(models.TableNameFriendRelation).Where("status=0 and user_id=?", uctx.UID(h.Ctx)).Find(&friendList).Error
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
			h.FriendUserInfoMap[user.UID].Name = user.UserName
		}
		h.FriendUserInfoMap[user.UID].Avatar = utils.GetPic(user.Avatar)
		h.FriendUserInfoMap[user.UID].Pinyin = utils.GetFirstPinYin(h.FriendUserInfoMap[user.UID].Name)
	}
}

func (h *GetFriendListHandler) PackData() {
	for _, friendInfo := range h.FriendUserInfoMap {
		h.Resp.UserList = append(h.Resp.UserList, friendInfo)
	}
	// 按首字符排序
	sort.Slice(h.Resp.UserList, func(i, j int) bool {
		return h.Resp.UserList[i].Pinyin < h.Resp.UserList[j].Pinyin
	})
	h.Resp.Count = int64(len(h.Resp.UserList))
}
