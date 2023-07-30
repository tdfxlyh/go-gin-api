package cronloader

import (
	"context"
	"github.com/jasonlvhit/gocron"
	"github.com/tdfxlyh/go-gin-api/dal/models"
)

// 定时读取的数据可以放到内存里
var (
	ctx          context.Context
	UserInfoList []*models.User
)

func InitCronLoader() {
	var err error
	ctx = context.Background()

	err = getUserInfoList()
	if err != nil {
		panic(err)
	}
	gocron.Every(180).Seconds().Do(getUserInfoList)
	go func() {
		<-gocron.Start()
	}()
}
