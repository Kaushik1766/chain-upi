package wallet

import (
	"fmt"
	"net/http"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/gin-gonic/gin"
)

type reqForm struct {
	Address string `json:"address" binding:"required"`
	Chain   string `json:"chain" binding:"required"`
}

func SetPrimary() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req reqForm
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
		}
		uid, exists := ctx.Get("uid")
		if !exists {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		fmt.Println(uid)
		err := db.SetPrimary(req.Address, uid.(string), req.Chain)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.AbortWithStatus(http.StatusOK)
	}
}
