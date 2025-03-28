package transaction

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/eth"
	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/trx"
	"github.com/gin-gonic/gin"
)

type transactionUPIForm struct {
	Amount       float64 `json:"amount" binding:"required"`
	ReceiverUPI  string  `json:"receiverUpi" binding:"required"`
	Chain        string  `json:"chain" binding:"required"`
	SenderWallet *string `json:"wallet"`
}

func SendToUpi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var form transactionUPIForm
		if err := ctx.ShouldBindJSON(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		senderUid, exists := ctx.Get("uid")
		if !exists {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		var senderWallet *models.Wallet
		var receiverWallet *models.Wallet

		// case - sender does not provides a wallet for payment
		if form.SenderWallet == nil {
			fmt.Println("sender wallet not provided")
			ctx.Status(200)
			return
		}

		// if he provides a wallet for payment
		senderWallet, err := db.GetUserWallet(senderUid.(string), *form.SenderWallet, form.Chain)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wallet not associated with user"})
			return
		}

		receiverWallet, err = db.GetPrimaryWalletByUpiHandle(form.ReceiverUPI, form.Chain)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error with receiver upi"})
		}

		switch strings.ToLower(form.Chain) {
		case "trx":
			err = trx.SendTrx(senderWallet, receiverWallet.Address, form.Amount)
		case "eth":
			err = eth.SendEth(ctx, senderWallet, receiverWallet.Address, form.Amount)
		default:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid chain"})
			return
		}
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

	}
}

type transactionAddressForm struct {
	Amount          float64 `json:"amount" binding:"required"`
	ReceiverAddress string  `json:"receiverAddress" binding:"required"`
	Chain           string  `json:"chain" binding:"required"`
	SenderWallet    string  `json:"wallet"`
}

func SendToAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
