package transaction

import (
	"net/http"

	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/trx"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	WalletAddress string `json:"wallet_address"`
}

func TransactionHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody RequestBody

		// Bind JSON request
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		transactions := trx.GetTransactions(reqBody.WalletAddress)

		if transactions.Status != "1" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": transactions.Message})
			return
		}

		c.JSON(http.StatusOK, gin.H{"transactions": transactions.Result})
	}
}
