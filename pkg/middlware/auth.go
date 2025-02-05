package middlware

import (
	"fmt"
	"net/http"

	"github.com/Kaushik1766/chain-upi-gin/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		token = token[7:]
		parsedToken, err := utils.ValidateJwt(token)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Println(parsedToken.Claims)
		ctx.Next()
	}
}
