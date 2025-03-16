package transaction

import (
	"net/http"
	"strings"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/eth"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	WalletAddress string `json:"walletAddress"`
	Chain         string `json:"chain"`
}

func TransactionHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody RequestBody
		// fmt.Println(c.Query("hello"))
		// Bind JSON request
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		var transactions []models.Transaction
		var err error
		switch strings.ToLower(reqBody.Chain) {
		case "eth":
			transactions, err = eth.GetTransactions(reqBody.WalletAddress)
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Error fetching transactions",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"transactions": transactions})
	}
}
