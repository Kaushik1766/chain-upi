package profile

import (
	"fmt"
	"net/http"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type changePasswordForm struct {
	Password string `json:"password" binding:"required"`
}

func ChangePassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, ok := ctx.Get("uid")
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		var req changePasswordForm
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid data",
			})
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		err = db.UpdatePassword(uid.(string), string(hashedPassword))
		if err != nil {
			fmt.Println(err.Error())
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.Status(http.StatusOK)
	}
}
