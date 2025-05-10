package transaction

import (
	"net/http"
	"strings"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/eth"
	"github.com/Kaushik1766/chain-upi-gin/pkg/crypto/trx"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	WalletAddress string `form:"walletAddress"`
	Chain         string `form:"chain"`
	UpiID         string `form:"upi"`
}

// type QueryBody struct {
// 	WalletAddress string `form:"walletAddress"`
// 	Chain         string `form:"chain"`
// 	UpiID		  string `form:"upi_id"`
// }

func TransactionHistoryByAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody RequestBody
		// var queryBody QueryBody
		// fmt.Println(c.Query("hello"))
		if err := c.ShouldBindQuery(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Query"})
			return
		}
		// queryBody.wallet_address = c.Query("WalletAddress")
		// queryBody.chain = c.Query("Chain")
		var transactions []models.Transaction
		var err error
		switch strings.ToLower(reqBody.Chain) {
		case "eth":
			transactions, err = eth.GetTransactions(reqBody.WalletAddress)
		case "trx":
			transactions, err = trx.GetTransactions(reqBody.WalletAddress)
		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Chain"})
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

func TransactionHistoryByUpi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody RequestBody
		if err := ctx.ShouldBindQuery(&reqBody); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Query"})
		}

		wallets, err := db.GetPrimaryWalletsByUpiHandle(reqBody.UpiID)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error fetching wallets"})
			return
		}

		for _, val := range wallets {
			var transactions []models.Transaction
			var err error
			switch strings.ToLower(val.Chain) {
			case "eth":
				transactions, err = eth.GetTransactions(val.Address)
			case "trx":
				transactions, err = trx.GetTransactions(val.Address)
			default:
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Chain"})
			}
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": "Error fetching transactions",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
		}
	}
}
