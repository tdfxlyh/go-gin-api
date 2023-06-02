package cronloader

import (
	"fmt"
	"github.com/tdfxlyh/go-gin-api/dal/models"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/utils"
	"log"
)

func getUserInfoList() error {
	userInfoList := make([]*models.User, 0)
	caller.LyhTestDB.Debug().Table("user").Find(&userInfoList)
	if userInfoList == nil || len(userInfoList) == 0 {
		log.Println("[GetUserInfoList] userInfoList is nil.")
		err := fmt.Errorf("userInfoList is nil")
		return err
	}
	log.Printf("[GetUserInfoList] userInfoList=%s", utils.GetStuStr(userInfoList))
	UserInfoList = userInfoList
	return nil
}
