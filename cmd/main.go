package main

import (
	"fmt"
	"log"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/Kaushik1766/chain-upi-gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cant find .env")
	}
	db, err := db.InitDB()
	// db.AutoMigrate(&models.User{}, &models.Wallet{})
	if db != nil {
		fmt.Println("DB connected.")
	}
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	v1 := r.Group("api/")
	routes.CreateRoutes(v1)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":3000")
}
