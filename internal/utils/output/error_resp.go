package output

type ErrorResp struct {
	Prompts   string `json:"prompts"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
	CustomMsg string `json:"custom_msg"`
}

func NewErrorResp(errCode IErrorCode, customMsg string) (resp *ErrorResp) {
	return &ErrorResp{
		Prompts:   errCode.Prompts(),
		Status:    errCode.Status(),
		Message:   errCode.Message(),
		CustomMsg: customMsg,
	}
}
