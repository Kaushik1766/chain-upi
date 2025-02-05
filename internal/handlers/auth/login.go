package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type loginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody loginForm
		_secret := []byte(os.Getenv("SECRET"))
		if err := ctx.ShouldBind(&reqBody); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Data",
			})
		}
		unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtClaims{
			Email: "kaushiksaha@gmail.com",
			Name:  reqBody.Username,
			UID:   "fasdf",
		})
		token, err := unsignedToken.SignedString(_secret)
		if err != nil {
			fmt.Println(err)
		}
		ctx.SetCookie("token", token, 3600, "/", ctx.Request.Host, false, true)
		fmt.Println(token)
		fmt.Println(reqBody.Username)
		fmt.Println("reqBody = ", reqBody)

	}
}
