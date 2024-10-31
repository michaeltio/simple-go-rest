package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID         int		`json:"id"`
	Item       string	`json:"item"`
	IsComplete bool		`json:"isComplete"`
}

var todos = []Todo{
	{ID: 1, Item: "Learn Go", IsComplete: false},
	{ID: 2, Item: "Make to do list", IsComplete: true},
	{ID: 3, Item: "Make database", IsComplete: false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoByID(id string) (*Todo, error) {
	// Konversi ID dari string ke int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	// Cari todo berdasarkan ID
	for i, t := range todos {
		if t.ID == idInt {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context){
	id := context.Param("id")
	todo, err:= getTodoByID(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return 
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func addTodo(context *gin.Context){
	var newTodo Todo

	if err := context.BindJSON(&newTodo); err != nil{
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated,newTodo)
}

func main(){
	router:= gin.Default();
	router.GET("todos",getTodos);
	router.GET("/todos/:id",getTodo);
	router.POST("todos",addTodo);
	router.Run("localhost:9090");
}