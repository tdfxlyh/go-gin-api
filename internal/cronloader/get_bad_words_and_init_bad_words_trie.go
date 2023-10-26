package cronloader

import (
	"fmt"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"log"
)

func getBadWordsAndInitBadWordsTrie() error {
	// 从数据库获取信息
	badWordsList, err := getBadWordsFromDB()
	if err != nil {
		return err
	}
	// 初始化敏感词树
	BadWordsTrie = initTrie(badWordsList)
	return nil
}

func getBadWordsFromDB() ([]string, error) {
	badWords := make([]*models.BadWord, 0)
	caller.LyhTestDB.Debug().Table(models.TableNameBadWord).Find(&badWords)
	if badWords == nil {
		log.Println("[getBadWordsFromDB] badWords is nil.")
		return nil, fmt.Errorf("badWords is nil")
	}
	badWordsList := make([]string, 0)
	for _, item := range badWords {
		badWordsList = append(badWordsList, item.Content)
	}
	log.Printf("[getBadWordsList] badWordsList=%v", badWordsList)
	return badWordsList, nil
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
