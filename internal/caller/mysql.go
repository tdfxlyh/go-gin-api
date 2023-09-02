package caller

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	LyhTestDB *gorm.DB
)

func InitDB() {
	LyhTestDB = InitDBDetail("root", "991113", "lyhtest", "127.0.0.1", 3306)
	//LyhTestDB = InitDBDetail("root", "991113", "lyhtest", "192.168.23.128", 3306)

	fmt.Println("success...")
	return
}

func InitDBDetail(username, password, dbname, ip string, port int64) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, ip, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
