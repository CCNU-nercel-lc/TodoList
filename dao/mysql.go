package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// DB 定义全局变量
var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	return DB.DB().Ping()
}
