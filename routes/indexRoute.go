package routes

import (
	"simple-go-rest/controllers"

	"github.com/gin-gonic/gin"
)

// index routes
func SetupIndexRoutes(router *gin.Engine) {
    router.GET("/", controllers.GetIndex)
}