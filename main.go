package main

import (
	"github.com/gin-gonic/gin"
	"golern/common"
	"golern/routers"
)

func main() {
	common.InitDb()
	r := gin.Default()
	r = routers.CollectRoutes(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
