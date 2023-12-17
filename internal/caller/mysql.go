package caller

import (
	"fmt"
	"github.com/tdfxlyh/go-gin-api/internal/constdef"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	LyhTestDB = InitDBDetail(constdef.DBLyhTestUserName, constdef.DBLyhTestPassword, constdef.DBLyhTestDBName, constdef.DBLyhTestIp, constdef.DBLyhTestPort)

	Logger.Info("db init success...")
	return
}

func InitDBDetail(username, password, dbname, ip string, port int64) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, ip, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Error("db init fail...")
		panic(err)
	}
	return db
}
