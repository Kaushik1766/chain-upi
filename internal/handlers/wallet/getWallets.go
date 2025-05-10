package wallet

import (
	"net/http"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/gin-gonic/gin"
)

func GetWallets() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, ok := ctx.Get("uid")
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		wallets, err := db.GetWalletsByUid(uid.(string))
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var resp []map[string]interface{}

		for _, val := range wallets {
			resp = append(resp, map[string]interface{}{
				"address":   val.Address,
				"chain":     val.Chain,
				"isPrimary": val.IsPrimary,
			})
		}

		ctx.AbortWithStatusJSON(http.StatusOK, resp)
	}
}
