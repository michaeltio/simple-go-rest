package utils

import (
	"fmt"
	"log"
	"os"
	"simple-go-rest/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    fmt.Printf("Connecting to database\n");
    
    //load env file from .env
    fmt.Printf("Load Env File\n");
    errenv := godotenv.Load()

    if( errenv != nil){
        log.Fatal("Error loading .env file")
    }
    

    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("POSTGRES_USER")
    dbPassword := os.Getenv("POSTGRES_PASSWORD")
    dbName := os.Getenv("POSTGRES_DB")
    dbPort := os.Getenv("DB_PORT")

    //connecting to postgres
    fmt.Printf("Connecting to Database");
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPassword, dbName, dbPort)
    fmt.Printf("DSN: %s\n", dsn);
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB.AutoMigrate(&models.Todo{})
}