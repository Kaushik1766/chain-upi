package routes

import (
	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/auth"
	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/profile"
	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/wallet"
	"github.com/Kaushik1766/chain-upi-gin/pkg/middlware"
	"github.com/gin-gonic/gin"
)

func CreateRoutes(r *gin.RouterGroup) {
	// authGroup := r.Group("auth/", middlware.Authenticate())
	authGroup := r.Group("/auth")
	authGroup.POST("/login", auth.Login())
	authGroup.POST("/signup", auth.Signup())

	profileGroup := r.Group("/profile", middlware.Authenticate())
	profileGroup.POST("/changePassword", profile.ChangePassword())

	walletGroup := r.Group("/wallet", middlware.Authenticate())
	walletGroup.POST("/addWallet", wallet.AddWallet())
}
