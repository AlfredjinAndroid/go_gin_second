package common

import (
	"fmt"
	"go_gin_second/model"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

//DB 数据库操作
var DB *gorm.DB

//InitDb 初始化数据库配置
func InitDb() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	userName := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		userName, password, host, port, database, charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("连接数据库失败 : " + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

//GetDB 获取数据库操作对象
func GetDB() *gorm.DB {
	return DB
}
