package dto_message

import "io/fs"

type AddMessageReq struct {
	MessageType    int64   `json:"message_type"`
	Content        string  `json:"content"`
	ReceiverUserID int64   `json:"receiver_user_id"`
	Timestamp      int64   `json:"timestamp"`
	File           fs.File `json:"file"`
}

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

type OptMessageReq struct {
	OptType   int64 `json:"opt_type"`
	MessageID int64 `json:"message_id"`
}
