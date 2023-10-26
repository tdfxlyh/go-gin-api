package utils

import (
	"github.com/tdfxlyh/go-gin-api/internal/cronloader"
	"strings"
)

func CheckBadWords(message string) bool {
	return isContain(message, cronloader.BadWordsTrie)
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
