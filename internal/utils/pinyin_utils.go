package utils

import (
	"github.com/mozillazg/go-pinyin"
	"unicode"
)

func GetFirstPinYin(str string) string {
	if str == "" {
		return ""
	}
	firstChar := rune(str[0])
	if unicode.Is(unicode.Han, firstChar) {
		// 如果是汉字
		pinyinSlice := pinyin.Pinyin(str, pinyin.NewArgs())
		return string(pinyinSlice[0][0][0])
	} else {
		// 不是汉字
		return SubString(str, 0, 1)
	}
}
