package routes

import (
	"fishing_company/pkg/controllers"
	"fishing_company/pkg/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func employeeRoutes(superRoute *gin.RouterGroup, e *casbin.Enforcer) {

	employeeRouter := superRoute.Group("/employees")
	employeeRouter.Use(middleware.AuthRequired)

	employeeRouter.GET("/", controllers.GetEmployees)
	employeeRouter.GET("/create", controllers.EmployeeForm)
	employeeRouter.POST("/create", controllers.CreateEmployee)
	employeeRouter.GET("/:id", controllers.GetEmployee)
	employeeRouter.GET("/:id/update", controllers.UpdateEmployeeForm)
	employeeRouter.POST("/:id/update", controllers.UpdateEmployee)
	employeeRouter.GET("/:id/delete", controllers.DeleteEmployeeForm)
	employeeRouter.DELETE("/:id/delete", controllers.DeleteEmployee)
}
