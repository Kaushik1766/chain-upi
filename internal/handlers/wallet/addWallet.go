package wallet

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/trx"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WalletForm struct {
	Address    string `json:"address" binding:"required"`
	PrivateKey string `json:"privateKey" binding:"required"`
	Chain      string `json:"chain" binding:"required"`
}

func AddWallet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var form WalletForm
		if err := ctx.ShouldBindJSON(&form); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		switch strings.ToLower(form.Chain) {
		case "trx":
			if !trx.ValidateAddress(form.Address) {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Chain"})
				return
			}
		}

		fmt.Println("Form: ", form)

		uid, ok := ctx.Get("uid")
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		parsedUid, err := uuid.Parse(uid.(string))
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		wallet := models.Wallet{
			Address:    form.Address,
			PrivateKey: form.PrivateKey,
			UserUID:    parsedUid,
			Chain:      form.Chain,
		}
		err = db.AddWallet(&wallet)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(200, gin.H{"message": "wallet added"})
	}
}
