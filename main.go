package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{"1", "Homework", true},
	{"2", "Task", false},
	{"3", "Cleaning", true},
}

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
func GetTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := GetTodoById(id)
	if err == nil {
		context.IndentedJSON(http.StatusOK, todo)
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}
}
func ToggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := GetTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}
func GetTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("there's no data with that id")
}
func AddTodo(context *gin.Context) {
	var newTodo todo
	err := context.Bind(&newTodo)
	if err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}
func main() {
	router := gin.Default()
	router.GET("/todos", GetTodos)
	router.GET("/todos/:id", GetTodo)
	router.PATCH("/todos/:id", ToggleTodoStatus)
	router.POST("/add", AddTodo)
	router.Run("localhost:9090")
}
