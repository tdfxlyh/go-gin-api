package dto_user

type OptUserReq struct {
	OptType  int64  `json:"opt_type"`
	Avatar   string `json:"avatar"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserInfoReq struct {
	OptType int64 `json:"opt_type"`
	UID     int64 `json:"uid"`
}

type UserInfoResp struct {
	UID      int64  `json:"uid"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
}
