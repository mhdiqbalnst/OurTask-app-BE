package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {

	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	// errEnv := godotenv.Load()
	// if errEnv != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	DB_USERNAME := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	var err error

	// config database

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=require password=%s port=%s", DB_HOST, DB_USERNAME, DB_NAME, DB_PASSWORD, DB_PORT)
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
