package repository

import (
	"batman/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ImageRepositoryInterface interface {
	Create(image *models.Image) error
	List(userID string) ([]models.Image, error)
	Get(imageID string) (*models.Image, error)
	Update(image *models.Image) error
	Delete(imageID string) error
}

type ImageRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewImageRepository(db *pgxpool.Pool) ImageRepositoryInterface {
	return &ImageRepositoryImpl{db: db}
}

func (r *ImageRepositoryImpl) Create(image *models.Image) error {
	_, err := r.db.Exec(context.Background(),
		"INSERT INTO image (image_id, user_id, original_filename, upload_date, width, height, file_size, file_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		image.ImageID, image.UserID, image.OriginalFileName, image.UploadDate, image.Width, image.Height, image.FileSize, image.FileType)
	fmt.Printf("Error from upload image repo %s", err)
	return err
}

func (r *ImageRepositoryImpl) List(userID string) ([]models.Image, error) {
	rows, err := r.db.Query(context.Background(),
		"SELECT image_id, user_id, original_filename, upload_date, width, height, file_size, file_type FROM image WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []models.Image
	for rows.Next() {
		var image models.Image
		if err := rows.Scan(&image.ImageID, &image.UserID, &image.OriginalFileName, &image.UploadDate, &image.Width, &image.Height, &image.FileSize, &image.FileType); err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	return images, nil
}

func (r *ImageRepositoryImpl) Get(imageID string) (*models.Image, error) {
	var image models.Image
	err := r.db.QueryRow(context.Background(),
		"SELECT image_id, user_id, original_filename, upload_date, width, height, file_size, file_type FROM image WHERE image_id = $1", imageID).
		Scan(&image.ImageID, &image.UserID, &image.OriginalFileName, &image.UploadDate, &image.Width, &image.Height, &image.FileSize, &image.FileType)
	return &image, err
}

func (r *ImageRepositoryImpl) Update(image *models.Image) error {
	query := "UPDATE image SET "
	params := []interface{}{}
	paramCounter := 1

	if image.UserID != "" {
		query += fmt.Sprintf("user_id = $%d, ", paramCounter)
		params = append(params, image.UserID)
		paramCounter++
	}
	if image.OriginalFileName != "" {
		query += fmt.Sprintf("original_filename = $%d, ", paramCounter)
		params = append(params, image.OriginalFileName)
		paramCounter++
	}
	if !image.UploadDate.IsZero() {
		query += fmt.Sprintf("upload_date = $%d, ", paramCounter)
		params = append(params, image.UploadDate)
		paramCounter++
	}
	if image.Width != 0 {
		query += fmt.Sprintf("width = $%d, ", paramCounter)
		params = append(params, image.Width)
		paramCounter++
	}
	if image.Height != 0 {
		query += fmt.Sprintf("height = $%d, ", paramCounter)
		params = append(params, image.Height)
		paramCounter++
	}
	if image.FileSize != 0 {
		query += fmt.Sprintf("file_size = $%d, ", paramCounter)
		params = append(params, image.FileSize)
		paramCounter++
	}
	if image.FileType != "" {
		query += fmt.Sprintf("file_type = $%d, ", paramCounter)
		params = append(params, image.FileType)
		paramCounter++
	}

	// Remove the last comma and space
	query = query[:len(query)-2]
	query += fmt.Sprintf(" WHERE image_id = $%d", paramCounter)
	params = append(params, image.ImageID)

	_, err := r.db.Exec(context.Background(), query, params...)
	return err
}

func (r *ImageRepositoryImpl) Delete(imageID string) error {
	_, err := r.db.Exec(context.Background(),
		"DELETE FROM image WHERE image_id = $1", imageID)
	return err
}
