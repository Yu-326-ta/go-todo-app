package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	// ユーザーの作成
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)
	// u.CreateUser()

	// ユーザーの取得
	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// // ユーザーの更新
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// ユーザーの削除
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// todoの作成
	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third Todo")

	// todoの取得
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// todosの取得
	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// userごとのtodos取得
	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// todoの更新
	// t, _ := models.GetTodo(1)
	// t.Content = "Update Todo"
	// t.UpdateTodo()

	t, _ := models.GetTodo(3)
	t.DeleteTodo()
}
