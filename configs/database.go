package configs

import (
    "fmt"
    "os"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
		"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Format DSN untuk PostgreSQL
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"))

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),})
    if err != nil {
        panic(fmt.Sprintf("Error initializing database: %v", err))
    } else {
        fmt.Println("Database connection successful")
    }
}
