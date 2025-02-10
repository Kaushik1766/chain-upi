package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type signupForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func Signup(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody signupForm
		// _secret := []byte(os.Getenv("SECRET"))
		if err := ctx.ShouldBind(&reqBody); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Data",
			})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 10)
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
			return
		}

		result := db.Create(&user)
		// fmt.Println(user.UID)
		if result.Error != nil {
			if strings.Contains(result.Error.Error(), "23505") {
				fmt.Println(result.Error.Error())
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "email address in use",
				})
				return
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Internal server error",
				})
				return
			}
		}
		ctx.JSON(http.StatusAccepted, gin.H{
			"success": "Account created Successfully, Login to continue.",
		})
	}
}
