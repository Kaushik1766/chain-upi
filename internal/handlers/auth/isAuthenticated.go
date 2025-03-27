package auth

import (
	"fmt"
	"net/http"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(200)
		return
	}
}

type checkPasswordReq struct {
	Password string `json:"password" binding:"required"`
}

func CheckPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody checkPasswordReq
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		uid, exists := ctx.Get("uid")
		if !exists {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := db.GetUserByUid(uid.(string))
		// fmt.Println(user.ToString())
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.AbortWithStatus(http.StatusOK)
	}
}
