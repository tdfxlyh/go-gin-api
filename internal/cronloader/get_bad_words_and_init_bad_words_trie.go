package cronloader

import (
	"fmt"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"log"
)

func getBadWordsAndInitBadWordsTrie() error {
	badWords := make([]*models.BadWord, 0)
	caller.LyhTestDB.Debug().Table(models.TableNameBadWord).Find(&badWords)
	if badWords == nil {
		log.Println("[getBadWordsList] badWords is nil.")
		err := fmt.Errorf("badWords is nil")
		return err
	}
	badWordsList := make([]string, 0)
	for _, item := range badWords {
		badWordsList = append(badWordsList, item.Content)
	}
	log.Printf("[getBadWordsList] badWordsList=%v", badWordsList)

	// 初始化敏感词树
	BadWordsTrie = initTrie(badWordsList)
	return nil
}

func initTrie(badWordsList []string) *BadWordsTrieNode {
	root := &BadWordsTrieNode{}
	for _, word := range badWordsList {
		addWord(root, word)
	}
	return root
}

type BadWordsTrieNode struct {
	Children map[rune]*BadWordsTrieNode
}

func addWord(root *BadWordsTrieNode, word string) {
	node := root
	for _, char := range word {
		if node.Children == nil {
			node.Children = make(map[rune]*BadWordsTrieNode)
		}
		if _, ok := node.Children[char]; !ok {
			node.Children[char] = &BadWordsTrieNode{}
		}
		node = node.Children[char]
	}
}
