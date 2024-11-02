package utils

import (
	"log"
	"simple-go-rest/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dsn := "host=localhost user=postgres password=root dbname=simple-go-rest port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB.AutoMigrate(&models.Todo{})
}