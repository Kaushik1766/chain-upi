package main

import (
	"fmt"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/trx"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Cant find .env")
	// }
	// err = db.InitDB()
	// // db.AutoMigrate(&models.User{}, &models.Wallet{})
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println("DB connected.")
	// }

	// parsedId, err := uuid.Parse("97bba45f-9ae7-40a7-81f8-fb0ec5e29eef")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// wallets, err := db.GetPrimaryWalletsByUpiHandle("kaushiksaha004")

	// err = db.SetPrimary("TCZVZUQpj1jB1YEoCMmhdAeNbCenkisrTq", parsedId, "trx")
	// wallets, err := db.GetWalletsByChain("kaushiksaha004", "trx")
	// if err != nil {
	// 	return
	// }
	// for _, val := range wallets {
	// 	fmt.Println(val.ToString())
	// }

	// wallet, err := db.GetPrimaryWalletByUpiHandle("kaushiksaha004", "trx")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(wallet.ToString())
	// for _, val := range wallets {
	// 	val.ToString()
	// }
	sender := models.Wallet{
		Address:    "TCZVZUQpj1jB1YEoCMmhdAeNbCenkisrTq",
		PrivateKey: "5D0BA2F54B2B3F329C2486BF831CDD8C6EBECF508FC1F704C6C4866F8A22FF52",
	}
	receiver := models.Wallet{
		Address: "TNqy91KcDjx5w9ZsmWzuPM8L8XJysbo19F",
	}
	err := trx.SendTrx(&sender, receiver.Address, 2000)
	if err != nil {
		fmt.Println(err)
	}

}
