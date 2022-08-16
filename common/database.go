package common

import (
	"fmt"
	"github.com/spf13/viper"
	"golern/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	errM := db.AutoMigrate(&model.User{}, &model.Document{})
	if errM != nil {
		panic("failed to migrate table:" + errM.Error())
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
