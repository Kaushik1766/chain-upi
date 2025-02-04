package routes

import (
	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/auth"
	"github.com/gin-gonic/gin"
)

func CreateRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("auth/")
	authGroup.POST("/login", auth.Login)
	authGroup.POST("/signup", auth.Signup)

}
