package transaction

import (
	"net/http"
	"strings"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/eth"
	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/trx"
	"github.com/gin-gonic/gin"
)

type transactionForm struct {
	Amount       float64 `json:"amount" binding:"required"`
	ReceiverUPI  string  `json:"receiverUpi" binding:"required"`
	Chain        string  `json:"chain" binding:"required"`
	SenderWallet string  `json:"sender"`
}

func SendToUpi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var form transactionForm
		if err := ctx.ShouldBindJSON(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		senderUpi, exists := ctx.Get("upi")
		if !exists {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		senderPrimaryWallet, err := db.GetPrimaryWalletByUpiHandle(senderUpi.(string), form.Chain)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Address not found for your account"})
			return
		}

		receiverPrimaryWallet, err := db.GetPrimaryWalletByUpiHandle(form.ReceiverUPI, form.Chain)

		if err != nil {
			// create a wallet for the user in the chain
		}

		var et error
		switch strings.ToLower(form.Chain) {
		case "trx":
			et = trx.SendTrx(senderPrimaryWallet, receiverPrimaryWallet.Address, form.Amount)
		case "eth":
			et = eth.SendEth(ctx, senderPrimaryWallet, receiverPrimaryWallet.Address, form.Amount)
		}
		if et != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": et.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Transaction created"})
	}
}
