package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/login", LoginForm)
	r.POST("/login", Login)
	r.GET("/logout", AuthRequired(), Logout)
	r.GET("/profile", AuthRequired(), TokenTimeoutRefresh(), Profile)
	r.POST("/register", Register)
	r.GET("/register", RegisterForm)
}
