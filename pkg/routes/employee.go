package routes

import (
	"fishing_company/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func employeeRoutes(superRoute *gin.RouterGroup) {

	employeeRouter := superRoute.Group("/employees")

	employeeRouter.GET("/", controllers.GetEmployees)
	employeeRouter.GET("/create", controllers.EmployeeForm)
	employeeRouter.POST("/create", controllers.CreateEmployee)
	employeeRouter.GET("/:id", controllers.GetEmployee)
	employeeRouter.GET("/:id/update", controllers.UpdateEmployeeForm)
	employeeRouter.POST("/:id/update", controllers.UpdateEmployee)
	employeeRouter.GET("/:id/delete", controllers.DeleteEmployeeForm)
	employeeRouter.POST("/:id/delete", controllers.DeleteEmployee)
}
