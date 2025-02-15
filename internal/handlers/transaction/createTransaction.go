package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionForm struct {
	Amount   float64 `json:"amount" binding:"required"`
	Receiver string  `json:"receiver" binding:"required"`
	Chain    string  `json:"chain" binding:"required"`
	Sender   string  `json:"sender" binding:"required"`
}

func CreateTransaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var form transactionForm
		if err := ctx.ShouldBindJSON(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		
		ctx.JSON(http.StatusOK, gin.H{"message": "Transaction created"})
	}
}
