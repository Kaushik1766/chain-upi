package wallet

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type balanceForm struct {
	Address string `json:"address" binding:"required"`
	Chain   string `json:"chain" binding:"required"`
}

func GetBalance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var form balanceForm
		if err := ctx.ShouldBindJSON(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		// switch strings.ToLower(form.Chain){
		// 	case ""
		// }
	}
}
