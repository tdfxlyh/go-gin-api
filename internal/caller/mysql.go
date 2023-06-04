package caller

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	LyhTestDB *gorm.DB
)

func InitDB() (err error) {
	dsn := "root:991113@tcp(127.0.0.1:3306)/lyhtest?charset=utf8&parseTime=True&loc=Local"
	LyhTestDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(fmt.Sprintf("database open err, err=%v", err))
		return
	}
	fmt.Println("success...")
	return
}
