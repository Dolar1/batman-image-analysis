package routes

import (
	"batman/app"
	handlers "batman/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(dependency *app.Dependency) *gin.Engine {
	router := gin.Default()
	router.GET("/ping", pingHandler)
	imageGroup := router.Group("/api/v1/images")
	{
		imageGroup.POST("/", handlers.UploadImageMetadata(dependency.ImageService))
		imageGroup.GET("/users/:user_id", handlers.ListImages(dependency.ImageService))
		imageGroup.GET("/:image_id", handlers.GetImageDetails(dependency.ImageService))
		imageGroup.PUT("/:image_id", handlers.UpdateImageMetadata(dependency.ImageService))
		imageGroup.DELETE("/:image_id", handlers.DeleteImage(dependency.ImageService))
	}
	return router
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
