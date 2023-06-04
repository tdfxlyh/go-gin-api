package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

// RandString 随机名字
func RandString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyz")
	result := make([]byte, n)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
