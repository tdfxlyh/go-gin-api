package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/constdef"
	"github.com/tdfxlyh/go-gin-api/internal/cronloader"
	"github.com/tdfxlyh/go-gin-api/internal/routers"
)

func main() {
	r := gin.Default()

	// 一些初始化操作
	caller.Init()
	// 定时任务初始化
	cronloader.InitCronLoader()

	r = routers.CollectRoute(r)

	r.Run(":" + constdef.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
