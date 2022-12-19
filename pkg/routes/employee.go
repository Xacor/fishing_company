package routes

import (
	"fishing_company/pkg/controllers"
	"fishing_company/pkg/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func employeeRoutes(superRoute *gin.RouterGroup, authEnforcer *casbin.Enforcer, isTesting bool) {

	employeeRouter := superRoute.Group("/employees")

	if !isTesting {
		employeeRouter.Use(middleware.AuthRequired, middleware.Authorization(authEnforcer))
	}

	employeeRouter.GET("/", controllers.GetEmployees)
	employeeRouter.GET("/create", controllers.EmployeeForm)
	employeeRouter.POST("/create", controllers.CreateEmployee)
	employeeRouter.GET("/:id", controllers.GetEmployee)
	employeeRouter.GET("/:id/update", controllers.UpdateEmployeeForm)
	employeeRouter.POST("/:id/update", controllers.UpdateEmployee)
	employeeRouter.POST("/:id/delete", controllers.DeleteEmployee)
}
