package utils

import (
	"batman/models"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RespondWithJSON is a utility function to send JSON responses
func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RespondWithError(c *gin.Context, success bool, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"success": success, "error": message})
}

// ValidateImageMetadata checks if the image metadata is valid
func ValidateImageMetadata(image *models.Image) error {
	if image.UserID == "" {
		return errors.New("user ID is required")
	}
	if image.OriginalFileName == "" {
		return errors.New("original name is required")
	}
	if image.Width <= 0 || image.Height <= 0 {
		return errors.New("invalid dimensions")
	}
	return nil
}
