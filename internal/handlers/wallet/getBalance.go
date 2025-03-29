package wallet

import (
	"net/http"
	"strings"

	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/trx"
	"github.com/gin-gonic/gin"
)

type balanceForm struct {
	Address string `json:"address" binding:"required"`
	Chain   string `json:"chain" binding:"required"`
}

func GetBalance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var form balanceForm
		if err := ctx.ShouldBindJSON(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		var balance float64
		var err error
		switch strings.ToLower(form.Chain) {
		case "trx":
			balance, err = trx.GetBalance(form.Address)
		case "eth":
			balance, err = trx.GetBalance(form.Address)
		default:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid chain"})
			return
		}
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting balance"})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"balance": balance})
	}
}
