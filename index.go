package main

import (
	"simple-go-rest/routes"
	"simple-go-rest/utils"

	"github.com/gin-gonic/gin"
)

func main(){
	utils.Connect();

	router:= gin.Default();
	routes.SetupRoutes(router);
	router.Run("localhost:8080");
}