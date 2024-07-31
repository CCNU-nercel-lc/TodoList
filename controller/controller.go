package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mybubble/dao"
	"mybubble/models"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo) // 接受前端的数据，赋值给todo

	// 将todo存入数据库
	if err := dao.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	if todoList, err := models.GetAllTodo(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error:": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(c *gin.Context) {
	// 从前端拿到id参数
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	// 到数据库中验证id是否存在
	//若不存在
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err:": err.Error()})
		return
	}

	// 若存在，正确从数据库拿到数据
	// 接受前端的数据，赋值给todo
	c.BindJSON(todo)
	// 将todo保存到数据库
	if err := models.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	fmt.Printf("第%v条数据修改完成\n", id)
}

func DeleteTodo(c *gin.Context) {
	// 从前端拿到id参数
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	//若删除失败
	if err := models.DeleteTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"err:": err.Error()})
		fmt.Printf("error: %v", err)
		return
	} else { // 删除成功
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}

	fmt.Printf("第%v条数据删除完成\n", id)

}
