package caller

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger    *zap.Logger // 日志
	LyhTestDB *gorm.DB    // 数据库
)

func Init() {
	// 日志设置
	InitLogger()

	// 初始化数据库
	InitDB()
}
