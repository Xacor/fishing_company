package routes

import (
	"github.com/Xacor/fishing_company/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func authRoutes(superRoute *gin.RouterGroup) {

	authRouter := superRoute.Group("/auth")
	authRouter.GET("/login", controllers.LoginForm)
	authRouter.POST("/login", controllers.Login)

	authRouter.GET("/register", controllers.RegisterForm)
	authRouter.POST("/register", controllers.Register)

	authRouter.GET("/logout", controllers.Logout)
}
