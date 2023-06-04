package caller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Init() {
	// 初始化数据库
	if err := InitDB(); err != nil {
		fmt.Println(fmt.Sprintf("database err, err=%v", err))
	}

	// 日志设置
	InitLogger()
}

func InitLogger() {
	// 配置将日志打印到文件内,---需要了解一下怎么写入
	file, _ := os.Create("./go-gin-api.log")
	gin.DefaultWriter = io.MultiWriter(file)
}
