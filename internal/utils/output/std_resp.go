package output

import "reflect"

type StdResp struct {
	Prompts string      `json:"prompts"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewStdResp(data interface{}) (resp *StdResp) {
	// data 为空时用 {}
	if data == nil || IsNil(data) {
		data = map[string]string{}
	}

	resp = &StdResp{
		Prompts: StatusCodeSuccess.Prompts(),
		Status:  StatusCodeSuccess.Status(),
		Message: StatusCodeSuccess.Message(),
		Data:    data,
	}

	return
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
