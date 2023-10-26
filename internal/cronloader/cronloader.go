package cronloader

import (
	"context"
	"github.com/jasonlvhit/gocron"
)

// 定时读取的数据可以放到内存里
var (
	ctx          context.Context
	BadWordsTrie *BadWordsTrieNode
)

func InitCronLoader() {
	var err error
	ctx = context.Background()

	err = getBadWordsAndInitBadWordsTrie() // 初始化敏感词树
	if err != nil {
		panic(err)
	}
	gocron.Every(3).Minutes().Do(getBadWordsAndInitBadWordsTrie)

	go func() {
		<-gocron.Start()
	}()
}
