package middlware

import (
	"net/http"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/gin-gonic/gin"
)

func Verify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		upi, _ := ctx.Get("uid")
		walletAddress := ctx.Query("walletAddress")
		err := db.VerifyWallet(walletAddress, upi.(string))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		ctx.Next()
	}
}
