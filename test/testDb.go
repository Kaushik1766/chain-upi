package main

import (
	"fmt"
	"log"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cant find .env")
	}
	err = db.InitDB()
	// db.AutoMigrate(&models.User{}, &models.Wallet{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB connected.")
	}

	parsedId, err := uuid.Parse("97bba45f-9ae7-40a7-81f8-fb0ec5e29eef")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = db.SetPrimary("TCZVZUQpj1jB1YEoCMmhdAeNbCenkisrTq", parsedId, "trx")
	// wallets, err := db.GetWalletsByChain("kaushiksaha004", "trx")
	// if err != nil {
	// 	return
	// }
	// for _, val := range wallets {
	// 	fmt.Println(val.ToString())
	// }
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
