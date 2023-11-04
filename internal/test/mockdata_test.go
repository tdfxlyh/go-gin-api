package test

import (
	"fmt"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"testing"
)

func TestMock(t *testing.T) {
	// 一些初始化操作
	caller.Init()
	fmt.Printf("init finish\n")

	sta := 1
	end := 200
	for i := sta; i <= end; i++ {
		if i%1000 == 0 {
			fmt.Printf("%.2f\n", float64(i)/float64(end))
		}
		res1 := &models.MessageSingle{
			SenderUserID:   1,
			ReceiverUserID: 3,
			MessageType:    1,
			Content:        fmt.Sprintf("你好呀_%d", i),
		}
		res2 := &models.MessageSingle{
			SenderUserID:   3,
			ReceiverUserID: 1,
			MessageType:    1,
			Content:        fmt.Sprintf("我很好，你也好呀_%d", i),
		}
		caller.LyhTestDB.Create(res1)
		caller.LyhTestDB.Create(res2)
	}
}
