package routes

import (
	"github.com/Kaushik1766/chain-upi-gin/internal/handlers/auth"
	"github.com/Kaushik1766/chain-upi-gin/pkg/middlware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRoutes(r *gin.RouterGroup, db *gorm.DB) {
	// authGroup := r.Group("auth/", middlware.Authenticate())
	authGroup := r.Group("/auth")
	authGroup.POST("/login", auth.Login(db))
	authGroup.POST("/signup", middlware.Authenticate(), auth.Signup())

}
