package controllers

import (
	"errors"
	"net/http"
	"simple-go-rest/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{
	{ID: 1, Item: "Learn Go", IsComplete: false},
	{ID: 2, Item: "Make to do list", IsComplete: true},
	{ID: 3, Item: "Make database", IsComplete: false},
}

func GetTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoByID(id string) (*models.Todo, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	for i, t := range todos {
		if t.ID == idInt {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func GetTodo(context *gin.Context){
	id := context.Param("id")
	todo, err:= getTodoByID(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return 
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func AddTodo(context *gin.Context){
	var newTodo models.Todo

	if err := context.BindJSON(&newTodo); err != nil{
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated,newTodo)
}

func ToggleTodoStatus(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoByID(id)

	if(err != nil){
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	todo.IsComplete = !todo.IsComplete

	context.IndentedJSON(http.StatusOK, todo)
}

func DeleteTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoByID(id)

	if(err != nil){
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	for i, t := range todos {
        if t.ID == todo.ID {
            todos = append(todos[:i], todos[i+1:]...)
            break
        }
    }

    context.IndentedJSON(http.StatusOK, gin.H{"message": "todo deleted"})
}