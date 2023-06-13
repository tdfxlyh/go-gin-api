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

func Int64Ptr(val int64) *int64 {
	return &val
}

func Int32Ptr(val int32) *int32 {
	return &val
}

func IntPtr(val int) *int {
	return &val
}

func Float64Ptr(val float64) *float64 {
	return &val
}

func Float32Ptr(val float32) *float32 {
	return &val
}

func StringPtr(val string) *string {
	return &val
}
