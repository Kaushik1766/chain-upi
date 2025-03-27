package wallet

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type reqForm struct {
	Address string `json:"address" binding:"required"`
	Chain   string `json:"chain" binding:"required"`
}

func SetPrimary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req reqForm
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
		}
	}
}
