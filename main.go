package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tdfxlyh/go-gin-api/internal/caller"
	"github.com/tdfxlyh/go-gin-api/internal/constdef"
	"github.com/tdfxlyh/go-gin-api/internal/routers"
)

func main() {
	r := gin.Default()

	caller.Init()

	r = routers.CollectRoute(r)

	r.Run(":" + constdef.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
