package main

import (
	"simple-go-rest/routes"
	"simple-go-rest/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    utils.Connect();

    router := gin.Default();

	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
				AllowHeaders: []string{"*"},
			}))

    //routes
    routes.SetupIndexRoutes(router);
    routes.SetupTodoRoutes(router);

    router.Run(":8080");
}