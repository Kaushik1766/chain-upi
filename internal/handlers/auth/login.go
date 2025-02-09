package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type loginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody loginForm
		_secret := []byte(os.Getenv("SECRET"))
		if err := ctx.ShouldBind(&reqBody); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Data",
			})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 10)
		fmt.Println(string(hashedPassword))
		user := models.User{
			Email:     reqBody.Email,
			UpiHandle: strings.Split(reqBody.Email, "@")[0],
			Name:      reqBody.Username,
			Password:  string(hashedPassword),
		}
		if err != nil {
			fmt.Println(err.Error())
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		result := db.Create(&user)
		if result.Error != nil {
			fmt.Println(result.Error.Error())
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
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
