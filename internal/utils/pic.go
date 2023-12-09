package utils

import (
	"net/http"
)

// GetPic 待扩展,未来可添加前缀等逻辑
func GetPic(pic string) string {

	return pic
}

// CheckOnlinePic 检查是否网络可访问
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
