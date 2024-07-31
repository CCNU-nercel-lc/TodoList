package routers

import (
	"github.com/gin-gonic/gin"
	"mybubble/controller"
)

func SetupRouters() *gin.Engine {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 代办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改某一个
		v1Group.PUT("todo/:id", controller.UpdateTodo)
		// 删除某一个
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	return r
}
