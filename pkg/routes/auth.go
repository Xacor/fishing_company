package routes

import (
	"fishing_company/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func authRoutes(superRoute *gin.RouterGroup) {

	authRouter := superRoute.Group("/auth")

	authRouter.GET("/login", controllers.LoginForm)
	authRouter.POST("/login", controllers.Login)
	authRouter.GET("/logout", controllers.AuthRequired(), controllers.Logout)
	authRouter.GET("/profile", controllers.AuthRequired(), controllers.TokenTimeoutRefresh(), controllers.Profile)
	authRouter.POST("/register", controllers.Register)
	authRouter.GET("/register", controllers.RegisterForm)
}
