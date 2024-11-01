package main

import (
	"simple-go-rest/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	router:= gin.Default();
	routes.SetupRoutes(router);
	router.Run("localhost:9090");
}