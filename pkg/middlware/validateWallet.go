package middlware

import (
	"net/http"
	"strings"

	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/wallet"
	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/trx"
	"github.com/gin-gonic/gin"
)

func ValidateWallet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody wallet.WalletForm
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		switch strings.ToLower(reqBody.Chain) {
		case "trx":
			if trx.ValidateAddress(reqBody.Address) {
				ctx.Next()
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Address"})
			}
		}
	}
}
