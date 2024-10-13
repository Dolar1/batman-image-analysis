package image

import (
	"batman/models"
	"batman/repository"
	"time"

	"github.com/google/uuid"
)

type ImageService interface {
	CreateImage(image *models.Image) error
	ListImages(userID string) ([]models.Image, error)
	GetImage(imageID string) (*models.Image, error)
	UpdateImage(image *models.Image) error
	DeleteImage(imageID string) error
}

type ImageServiceImpl struct {
	repo repository.ImageRepositoryInterface
}

func NewImageService(repo repository.ImageRepositoryInterface) ImageService {
	return &ImageServiceImpl{repo: repo}
}

func (s *ImageServiceImpl) CreateImage(image *models.Image) error {
	image.UploadDate = time.Now()
	image.ImageID = uuid.New().String()
	return s.repo.Create(image)
}

func (s *ImageServiceImpl) ListImages(userID string) ([]models.Image, error) {
	return s.repo.List(userID)
}

func (s *ImageServiceImpl) GetImage(imageID string) (*models.Image, error) {
	return s.repo.Get(imageID)
}

func (s *ImageServiceImpl) UpdateImage(image *models.Image) error {
	return s.repo.Update(image)
}

func (s *ImageServiceImpl) DeleteImage(imageID string) error {
	return s.repo.Delete(imageID)
}
