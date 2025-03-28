package transaction

import (
	"net/http"
	"strings"

	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/eth"
	"github.com/gin-gonic/gin"
)

func GetBalance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		chain := ctx.Query("chain")
		walletAddress := ctx.Query("walletAddress")

		if chain == "" || walletAddress == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing Parameters"})
			return
		}
		var bal string
		var err error
		switch strings.ToLower(chain) {
		case "eth":
			bal, err = eth.GetBalance(walletAddress)
		case "trx":
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Not Implemented"})
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Chain"})
		}
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{"balance": bal})
			return
		}
	}
}
