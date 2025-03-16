package routes

import (
	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/auth"
	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/profile"
	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/transaction"
	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/wallet"
	"github.com/Kaushik1766/chain-upi-gin/pkg/middlware"
	"github.com/gin-gonic/gin"
)

func CreateRoutes(r *gin.RouterGroup) {
	// authGroup := r.Group("auth/", middlware.Authenticate())
	authGroup := r.Group("/auth")
	authGroup.POST("/login", auth.Login())
	authGroup.POST("/signup", auth.Signup())
	// authGroup.POST("/test", Test())

	profileGroup := r.Group("/profile", middlware.Authenticate())
	profileGroup.POST("/changePassword", profile.ChangePassword())

	walletGroup := r.Group("/wallet", middlware.Authenticate())
	walletGroup.POST("/addWallet", middlware.ValidateWallet(), wallet.AddWallet())
	walletGroup.POST("/setPrimary", wallet.SetPrimary())

	transactionGroup := r.Group("/transaction", middlware.Authenticate())
	transactionGroup.POST("/sendToUpi", transaction.SendToUpi())

	transactionHistory := transactionGroup.Group("/history", middlware.Verify())
	transactionHistory.GET("/ByUpi", transaction.TransactionHistoryByUpi())
	transactionHistory.GET("/ByAddress", transaction.TransactionHistoryByAddress())

	// transactionGroup.POST("/sendToAddress", transaction.CreateTransaction())

}

// type testForm struct {
// 	// Address string `json:"address" binding:"required"`
// 	// PrivateKey string `json:"privateKey" binding:"required"`
// 	UPI   string `json:"upi" binding:"required"`
// 	Chain string `json:"chain" binding:"required"`
// }

// func Test() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var form testForm
// 		if err := ctx.ShouldBindJSON(&form); err != nil {
// 			ctx.JSON(400, gin.H{"error": "Invalid data"})
// 			return
// 		}
// 		_, err := db.GetPrimaryWalletByUpiHandle(form.UPI, form.Chain)
// 		if err != nil {
// 			ctx.JSON(400, gin.H{"error": "Invalid data2"})
// 			return
// 		}
// 		ctx.JSON(200, gin.H{"message": "wallet added"})
// 	}
// }
