package boats

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	routes := r.Group("/boats")
	routes.GET("/", GetBoats)
	routes.GET("/new", BoatForm)
	routes.POST("/create", CreateBoat)
	routes.GET("/:id", GetBoat)
}
