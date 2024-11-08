package controllers

import (
	"errors"
	"net/http"
	"simple-go-rest/models"
	"simple-go-rest/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//get all todo list
func GetTodos(context *gin.Context) {
    var todos []models.Todo
    utils.DB.Find(&todos)
    context.IndentedJSON(http.StatusOK, todos)
}

// get todo by id finder
func getTodoByID(id string) (*models.Todo, error) {
    var todo models.Todo
    idInt, err := strconv.Atoi(id)
    if err != nil {
        return nil, errors.New("invalid ID format")
    }

    if err := utils.DB.First(&todo, idInt).Error; err != nil {
        return nil, err
    }

    return &todo, nil
}

//get todo by id
func GetTodo(context *gin.Context) {
    id := context.Param("id")
    todo, err := getTodoByID(id)

    if err != nil {
        context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
        return
    }

    context.IndentedJSON(http.StatusOK, todo)
}

//add new todo list
func AddTodo(context *gin.Context) {
    var newTodo models.Todo

    if err := context.BindJSON(&newTodo); err != nil {
        return
    }

    utils.DB.Create(&newTodo)
    context.IndentedJSON(http.StatusCreated, gin.H{"message": "added successfully"})
}

//toggle todo status
func ToggleTodoStatus(context *gin.Context) {
    id := context.Param("id")
    todo, err := getTodoByID(id)

    if err != nil {
        context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
        return
    }

    todo.IsComplete = !todo.IsComplete
    utils.DB.Save(&todo)

    context.IndentedJSON(http.StatusOK, gin.H{"message": "status updated"})
}

//delete todo
func DeleteTodo(context *gin.Context) {
    id := context.Param("id")
    todo, err := getTodoByID(id)

    if err != nil {
        context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
        return
    }

    utils.DB.Delete(&todo)
    context.IndentedJSON(http.StatusOK, gin.H{"message": "todo deleted"})
}