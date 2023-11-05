package utils

import (
	"net/http"
)

func GetPic(pic string) string {

	return pic
}

func CheckOnlinePic(imageUrl string) bool {
	if imageUrl == "" {
		return false
	}
	// 发送HTTP请求
	response, err := http.Head(imageUrl)
	if err != nil {
		return false
	}
	// 检查响应状态码
	if response.StatusCode == http.StatusOK {
		return true
	}
	return false
}
