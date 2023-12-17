package cronloader

import (
	"strings"
)

// CheckBadWords 检查是否包含敏感词
func CheckBadWords(message string) bool {
	return isContain(message, BadWordsTrie)
}

// CheckAndReplaceBadWords 检查是否包含敏感词并替换为*
func CheckAndReplaceBadWords(message string) string {
	return replaceBadWords(message, '*', BadWordsTrie)
}

// CheckAndReplaceBadWordsWithSep 检查是否包含敏感词并替换为{sep}
func CheckAndReplaceBadWordsWithSep(message string, sep rune) string {
	return replaceBadWords(message, sep, BadWordsTrie)
}

func isContain(message string, root *BadWordsTrieNode) bool {
	message = strings.ToLower(message)
	runes := []rune(message)
	for i := 0; i < len(runes); i++ {
		p := root
		j := i
		for j < len(runes) && p.Children != nil {
			char := runes[j]
			if _, ok := p.Children[char]; ok {
				p = p.Children[char]
				j++
			} else {
				break
			}
		}
		if p.Children == nil {
			return true
		}
	}
	return false
}

func replaceBadWords(message string, sep rune, root *BadWordsTrieNode) string {
	message = strings.ToLower(message)
	runes := []rune(message)
	for i := 0; i < len(runes); i++ {
		p := root
		j := i
		for j < len(runes) && p.Children != nil {
			char := runes[j]
			if _, ok := p.Children[char]; ok {
				p = p.Children[char]
				j++
			} else {
				break
			}
		}
		if p.Children == nil {
			for k := i; k < j; k++ {
				runes[k] = sep
			}
		}
	}
	return string(runes)
}
