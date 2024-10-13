package handlers

import (
	image_model "batman/models"
	image_service "batman/services/image"
	"batman/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

// UploadImageMetadata handles uploading image metadata.
func UploadImageMetadata(service image_service.ImageService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var image image_model.Image
		if err := c.ShouldBindJSON(&image); err != nil {
			utils.RespondWithJSON(c.Writer, http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate image metadata
		if err := utils.ValidateImageMetadata(&image); err != nil {
			utils.RespondWithError(c, false, http.StatusBadRequest, err.Error())
			return
		}
		fmt.Printf("image is %v", image)
		if err := service.CreateImage(&image); err != nil {
			utils.RespondWithError(c, false, http.StatusInternalServerError, "Failed to upload image metadata")
			return
		}

		utils.RespondWithJSON(c.Writer, http.StatusOK, gin.H{"message": "Image metadata uploaded successfully", "image_id": image.ImageID})
	}
}

// ListImages handles retrieving all images for a specific user.
func ListImages(service image_service.ImageService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id") // Assuming you pass user_id as a path parameter
		images, err := service.ListImages(userID)
		if err != nil {
			utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, gin.H{"error": "Failed to retrieve images"})
			return
		}

		utils.RespondWithJSON(c.Writer, http.StatusOK, images)
	}
}

// GetImageDetails handles retrieving details for a specific image by its ID.
func GetImageDetails(service image_service.ImageService) gin.HandlerFunc {
	return func(c *gin.Context) {
		imageID := c.Param("image_id") // Assuming you pass image_id as a path parameter
		image, err := service.GetImage(imageID)
		if err != nil {
			if err == pgx.ErrNoRows {
				utils.RespondWithJSON(c.Writer, http.StatusNotFound, gin.H{"error": "Image not found"})
				return
			}
			utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, gin.H{"error": "Failed to retrieve image details"})
			return
		}

		utils.RespondWithJSON(c.Writer, http.StatusOK, image)
	}
}

// UpdateImageMetadata handles updating metadata for a specific image.
func UpdateImageMetadata(service image_service.ImageService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var image image_model.Image
		if err := c.ShouldBindJSON(&image); err != nil {
			utils.RespondWithJSON(c.Writer, http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		imageID := c.Param("image_id") // Assuming you pass image_id as a path parameter
		image.ImageID = imageID        // Set the image ID for the update

		if err := service.UpdateImage(&image); err != nil {
			utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, gin.H{"error": "Failed to update image metadata"})
			return
		}

		utils.RespondWithJSON(c.Writer, http.StatusOK, gin.H{"message": "Image metadata updated successfully"})
	}
}

// DeleteImage handles deleting a specific image by its ID.
func DeleteImage(service image_service.ImageService) gin.HandlerFunc {
	return func(c *gin.Context) {
		imageID := c.Param("image_id") // Assuming you pass image_id as a path parameter
		if err := service.DeleteImage(imageID); err != nil {
			if err == pgx.ErrNoRows {
				utils.RespondWithJSON(c.Writer, http.StatusNotFound, gin.H{"error": "Image not found"})
				return
			}
			utils.RespondWithJSON(c.Writer, http.StatusInternalServerError, gin.H{"error": "Failed to delete image"})
			return
		}

		utils.RespondWithJSON(c.Writer, http.StatusOK, gin.H{"message": "Image deleted successfully"})
	}
}
