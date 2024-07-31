package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"mybubble/dao"
	"mybubble/models"
	"mybubble/routers"
)

func main() {
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 延迟关闭数据库
	defer dao.DB.Close()

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{}) // 创建一个todos的表

	r := routers.SetupRouters()
	r.Run(":9090")
}
