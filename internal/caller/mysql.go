package caller

import (
	"fmt"
	"github.com/tdfxlyh/go-gin-api/internal/constdef"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	LyhTestDB *gorm.DB
)

func InitDB() {
	LyhTestDB = InitDBDetail(constdef.DBLyhTestUserName, constdef.DBLyhTestPassword, constdef.DBLyhTestDBName, constdef.DBLyhTestIp, constdef.DBLyhTestPort)

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
