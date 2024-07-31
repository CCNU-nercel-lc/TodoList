package models

import (
	"mybubble/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Todo 的增删改查

// GetAllTodo 查询所有
func GetAllTodo() (todoList []Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

// GetATodo 查询一条数据
func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo) //!!!这里不能省略
	if err := dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateTodo 修改
func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return err
}

// DeleteTodo 删除
func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(Todo{}).Error
	return
}
