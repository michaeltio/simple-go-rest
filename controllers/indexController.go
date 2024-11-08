package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//index routes
func GetIndex(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Welcome to the golang rest api",
    })
}