package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	_databaseUrl := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(_databaseUrl), &gorm.Config{})
	if err != nil {
		// log.Fatal("Failed to connect to database: ", err)
		return err
	}
	DB = db

	// reset db
	// DB.Delete(&models.User{}, "1=1")
	// DB.Delete(&models.Wallet{}, "1=1")

	// migrate if you change models
	// DB.Migrator().DropTable(&models.User{})
	// DB.Migrator().DropTable(&models.Wallet{})
	// DB.AutoMigrate(&models.User{})
	// DB.AutoMigrate(&models.Wallet{})
	fmt.Println("Connected to db.")
	return nil
}
