package utils

import (
	"encoding/json"
	"fmt"
)

// GetStuStr 格式化对象输出
func GetStuStr(obj interface{}) string {
	if obj == nil {
		return ""
	}
	bs, err := json.Marshal(obj)
	if err != nil {
		return fmt.Sprintf("Marshal Error. obj=%v err=%v", obj, err)
	}
	return string(bs)
}
