package routes

import (
	"evento-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/eventos", controllers.GetEventos)
		api.POST("/eventos", controllers.CreateEvento)
		api.GET("/eventos/:id", controllers.GetEventoByID)
		api.DELETE("/eventos/:id", controllers.DeleteEvento)
	}
}
