package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		// log.Fatal("No env file")
		return nil, err
	}
	_databaseUrl := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(_databaseUrl), &gorm.Config{})
	if err != nil {
		// log.Fatal("Failed to connect to database: ", err)
		return nil, err
	}
	fmt.Println("Connected to db.")
	return db, nil
}
