package utils

import (
	"github.com/mozillazg/go-pinyin"
	"unicode"
	"unicode/utf8"
)

func GetFirstPinYin(str string) string {
	if str == "" {
		return ""
	}
	firstChar, _ := utf8.DecodeRuneInString(str)
	if unicode.Is(unicode.Han, firstChar) {
		// 如果是汉字
		pinyinSlice := pinyin.Pinyin(str, pinyin.NewArgs())
		return string(pinyinSlice[0][0][0])
	} else {
		// 不是汉字
		return SubString(str, 0, 1)
	}
}
