package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	_databaseUrl := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(_databaseUrl), &gorm.Config{})
	if err != nil {
		// log.Fatal("Failed to connect to database: ", err)
		return nil, err
	}
	fmt.Println("Connected to db.")
	return db, nil
}
