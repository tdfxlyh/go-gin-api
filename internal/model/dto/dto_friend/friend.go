package dto_friend

type AddFriendReq struct {
	Phone string `json:"phone"`
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

type UpdateNotesReq struct {
	UserID int64  `json:"user_id"`
	Notes  string `json:"notes"`
}
