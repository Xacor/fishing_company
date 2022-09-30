package boats

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/boats")
	routes.GET("/", h.GetBoats)
	routes.GET("/new", h.BoatForm)
	routes.POST("/create", h.CreateBoat)
	routes.GET("/:id", h.GetBoat)
	routes.GET("/:id/update", h.UpdateBoatForm)
	routes.POST("/:id/update", h.UpdateBoat)
}
