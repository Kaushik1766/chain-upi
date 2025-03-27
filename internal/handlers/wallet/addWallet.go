package wallet

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/eth"
	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/trx"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WalletForm struct {
	PrivateKey string `json:"privateKey" binding:"required"`
	Chain      string `json:"chain" binding:"required"`
}

func AddWallet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var form WalletForm
		if err := ctx.ShouldBind(&form); err != nil {
			fmt.Println(err.Error())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid data."})
			return
		}
		var wallet *models.Wallet
		var err error
		switch strings.ToLower(form.Chain) {
		case "trx":
			wallet, err = trx.PrivateKeyToWallet(form.PrivateKey)
		case "eth":
			wallet, err = eth.PrivateKeyToWallet(form.PrivateKey)
		default:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "invalid chain",
			})
		}
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

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

		wallet.UserUID = parsedUid
		err = db.AddWallet(wallet)

		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "wallet added successfully",
			"wallet":  wallet.Address,
		})
	}
}
