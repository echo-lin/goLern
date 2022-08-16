package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golern/common"
	"golern/routers"
	"os"
)

func main() {
	InitConfig()
	common.InitDb()
	r := gin.Default()
	r = routers.CollectRoutes(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}

}

func InitConfig() {
	wordDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(wordDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
