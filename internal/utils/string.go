package utils

import "strings"

// SubString 包含中文的子字符串
func SubString(source string, start, end int) string {
	var unicodeStr = []rune(source)
	length := len(unicodeStr)
	if start >= end {
		return ""
	}
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	if start <= 0 && end >= length {
		return source
	}
	var substring = ""
	for i := start; i < end; i++ {
		substring += string(unicodeStr[i])
	}
	return substring
}

// StringLen 获取带中文的字符串实际长度
func StringLen(str string) int {
	var r = []rune(str)
	return len(r)
}

// StringIndex 包含中文的子字符串所在的位置
func StringIndex(str, subStr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, subStr)
	if result > 0 {
		prefix := []byte(str)[0:result]
		rs := []rune(string(prefix))
		result = len(rs)
	}
	return result
}
