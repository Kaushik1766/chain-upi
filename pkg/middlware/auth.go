package middlware

import (
	"fmt"
	"net/http"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/Kaushik1766/chain-upi-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if len(token) <= 7 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token = token[7:]
		parsedToken, err := utils.ValidateJwt(token)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Set("uid", parsedToken.UID)
		ctx.Set("upi", parsedToken.UpiHandle)
		ctx.Next()
	}
}

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
