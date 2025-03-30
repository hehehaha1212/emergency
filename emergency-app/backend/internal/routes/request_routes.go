package routes

import (
	"github.com/gin-gonic/gin"
	"emergency-app/internal/controllers"
	"emergency-app/internal/middleware"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		// Authenticated routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/request", controllers.CreateRequest)
			protected.GET("/request/:id", controllers.GetRequest)
			protected.PUT("/request/:id", controllers.UpdateRequest)
			protected.DELETE("/request/:id", controllers.DeleteRequest)
		}
	}
	
	AuthRoutes(router)
}