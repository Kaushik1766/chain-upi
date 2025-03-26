package routes

import (
	"time"

	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/auth"
	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/profile"
	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/transaction"
	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/wallet"
	"github.com/Kaushik1766/chain-upi-gin/pkg/middlware"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func CreateRoutes(r *gin.RouterGroup) {
	// authGroup := r.Group("auth/", middlware.Authenticate())
	authGroup := r.Group("/auth")
	authGroup.POST("/login", auth.Login())
	authGroup.POST("/signup", auth.Signup())
	authGroup.GET("/check", middlware.Authenticate(), auth.IsAuthenticated())
	// authGroup.POST("/test", Test())

	profileGroup := r.Group("/profile", middlware.Authenticate())
	profileGroup.POST("/changePassword", profile.ChangePassword())

	walletGroup := r.Group("/wallet", middlware.Authenticate())
	walletGroup.POST("/addWallet", wallet.AddWallet())
	walletGroup.POST("/setPrimary", wallet.SetPrimary())
	walletGroup.GET("/getWallets", wallet.GetWallets())

	transactionGroup := r.Group("/transaction", middlware.Authenticate())
	transactionGroup.POST("/sendToUpi", timeout.New(
		timeout.WithTimeout(10*time.Second),
		timeout.WithHandler(transaction.SendToUpi()),
	))

	transactionHistory := transactionGroup.Group("/history", middlware.Verify())
	transactionHistory.GET("/upi", transaction.TransactionHistoryByUpi())
	transactionHistory.GET("/address", transaction.TransactionHistoryByAddress())

	// transactionGroup.POST("/sendToAddress", transaction.CreateTransaction())

}
