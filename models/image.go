package models

import "time"

// Image represents the metadata of an uploaded image.
type Image struct {
	ImageID          string    `json:"image_id"`          // Unique identifier for the image
	UserID           string    `json:"user_id"`           // ID of the user who uploaded the image
	OriginalFileName string    `json:"original_filename"` // Original filename of the uploaded image
	UploadDate       time.Time `json:"upload_date"`       // Date and time when the image was uploaded
	Width            int       `json:"width"`             // Width of the image in pixels
	Height           int       `json:"height"`            // Height of the image in pixels
	FileSize         int64     `json:"file_size"`         // Size of the image file in bytes
	FileType         string    `json:"file_type"`         // MIME type of the image (e.g., image/jpeg)
}
