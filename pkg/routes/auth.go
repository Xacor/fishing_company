package routes

import (
	"fishing_company/pkg/controllers"
	"fishing_company/pkg/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func authRoutes(superRoute *gin.RouterGroup, e *casbin.Enforcer) {

	authRouter := superRoute.Group("/auth")
	authRouter.GET("/login", controllers.LoginForm)
	authRouter.POST("/login", controllers.Login)

	authRouter.GET("/register", controllers.RegisterForm)
	authRouter.POST("/register", controllers.Register)

	authRouter.GET("/logout", middleware.AuthRequired, controllers.Logout)

}
