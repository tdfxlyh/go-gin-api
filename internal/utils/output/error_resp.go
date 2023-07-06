package output

import (
	"net/http"
)

type ErrorResp struct {
	Prompts   string `json:"prompts"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
	CustomMsg string `json:"custom_msg"`
}

func (err *ErrorResp) HTTPCode() int {
	return http.StatusOK
}

func (err *ErrorResp) Errno() int {
	return err.Status
}

func (err *ErrorResp) ErrMsg() string {
	return err.Message
}

func (err *ErrorResp) ErrPrompts() interface{} {
	return err.Prompts
}

func (err *ErrorResp) Panic() interface{} {
	return nil
}

func NewErrorResp(errCode IErrorCode, customMsg string) (resp *ErrorResp) {
	resp = &ErrorResp{
		Prompts:   errCode.Prompts(),
		Status:    errCode.Status(),
		Message:   errCode.Message(),
		CustomMsg: customMsg,
	}

	return
}

type IErrorCode interface {
	Prompts() string
	Message() string
	Status() int
}

type ErrorCode int

const (
	StatusCodeSuccess        ErrorCode = 0
	StatusCodeParamsError    ErrorCode = 40001
	StatusCodeNoPermission   ErrorCode = 40002
	StatusCodeNotLoggedIn    ErrorCode = 40003
	StatusCodeSeverException ErrorCode = 50000
	StatusCodeDBError        ErrorCode = 50001
	StatusCodeRedisError     ErrorCode = 50002
	StatusCodeRPCError       ErrorCode = 50003
	StatusCodeTooBusy        ErrorCode = 50004
	StatusCodeImgDataFail    ErrorCode = 50005
	StatusCodeContentTooLong ErrorCode = 50006
)

func (e ErrorCode) Prompts() string {
	switch e {
	case StatusCodeSuccess:
		return ""
	case StatusCodeParamsError:
		return "请求参数错误"
	case StatusCodeNoPermission:
		return "无权限"
	case StatusCodeNotLoggedIn:
		return "未登录"
	case StatusCodeSeverException, StatusCodeDBError,
		StatusCodeRedisError, StatusCodeRPCError:
		return "服务器内部错误，请稍后重试"
	case StatusCodeTooBusy:
		return "操作太频繁了吧，请稍微休息下"
	case StatusCodeImgDataFail:
		return "图片转码失败"
	case StatusCodeContentTooLong:
		return "超过最大字数限制，请删减后再发布"
	}

	return "unknown error"
}

func (e ErrorCode) Message() string {
	switch e {
	case StatusCodeSuccess:
		return "success"
	case StatusCodeParamsError:
		return "params error"
	case StatusCodeNoPermission:
		return "no permission"
	case StatusCodeNotLoggedIn:
		return "not logged in"
	case StatusCodeSeverException, StatusCodeDBError,
		StatusCodeRedisError, StatusCodeRPCError:
		return "server exception"
	case StatusCodeTooBusy:
		return "too busy"
	case StatusCodeImgDataFail:
		return "image data decode fail"
	case StatusCodeContentTooLong:
		return "content too long"
	}

	return "unknown error"
}

func (e ErrorCode) Status() int {
	return int(e)
}
