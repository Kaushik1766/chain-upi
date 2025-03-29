package wallet

import (
	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/gin-gonic/gin"
)

type deleteForm struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
	Chain         string `json:"chain" binding:"required"`
}

func DeleteWallet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var form deleteForm
		if err := ctx.ShouldBindJSON(&form); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid data"})
			return
		}
		uid, exists := ctx.Get("uid")
		if !exists {
			ctx.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		err := db.DeleteWallet(form.WalletAddress, form.Chain, uid.(string))
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		ctx.JSON(200, gin.H{"message": "Wallet deleted successfully"})
	}
}
