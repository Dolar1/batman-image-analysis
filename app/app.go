package app

import (
	"batman/config"
	"batman/db"
	"batman/repository"
	"batman/services/image"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Dependency struct {
	DBPool       *pgxpool.Pool
	ImageService image.ImageService
}

func InitConfig() {
	config.Load()
	// log.SetupLogger()
}

func Init() (*Dependency, error) {
	InitConfig()

	dbPool, err := db.Connect(config.DBUrl())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to the database")
	}

	// Initialize the repositories
	imageRepo := repository.NewImageRepository(dbPool)

	// Initialize other services
	imageService := image.NewImageService(imageRepo)

	return &Dependency{
		DBPool:       dbPool,
		ImageService: imageService,
	}, nil
}

func StopApp() {
	fmt.Println("Stopping Batman Image Analysis Platform")
}
