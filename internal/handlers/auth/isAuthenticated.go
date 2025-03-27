package auth

import "github.com/gin-gonic/gin"

func IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(200)
		return
	}
}
