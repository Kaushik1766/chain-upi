package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type loginForm struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
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
		user, err := db.GetUser(reqBody.Email)
		// fmt.Println(user.ToString())
		if err != nil {
			fmt.Println(err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err": "Invalid credentials",
			})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// fmt.Println(user)
		unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtClaims{
			Email:     user.Email,
			Name:      user.Name,
			UID:       user.UID.String(),
			UpiHandle: user.UpiHandle,
		})
		token, err := unsignedToken.SignedString(_secret)
		if err != nil {
			fmt.Println(err)
		}
		ctx.SetCookie("token", token, 3600, "/", ctx.Request.Host, false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"success": "logged in",
		})
	}
}
