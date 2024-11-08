package routes

import (
	"simple-go-rest/controllers"

	"github.com/gin-gonic/gin"
)

//todo routes
func SetupTodoRoutes(router *gin.Engine) {
	router.GET("/todos", controllers.GetTodos);
	router.GET("/todos/:id", controllers.GetTodo);
	router.POST("todos", controllers.AddTodo);
	router.PATCH("/todos/:id", controllers.ToggleTodoStatus);
	router.DELETE("/todos/:id", controllers.DeleteTodo);
}