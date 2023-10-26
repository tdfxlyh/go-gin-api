package utils

import (
	"github.com/tdfxlyh/go-gin-api/internal/cronloader"
	"strings"
)

func CheckBadWords(message string) bool {
	return isContain(message, cronloader.BadWordsTrie)
}

func CheckAndReplaceBadWords(message string) string {
	return replaceBadWords(message, '*', cronloader.BadWordsTrie)
}

func CheckAndReplaceBadWordsWithSep(message string, sep rune) string {
	return replaceBadWords(message, sep, cronloader.BadWordsTrie)
}

func isContain(message string, root *cronloader.BadWordsTrieNode) bool {
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

func replaceBadWords(message string, sep rune, root *cronloader.BadWordsTrieNode) string {
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
