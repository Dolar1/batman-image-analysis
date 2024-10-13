package db

import (
	"context"
	"log"
	"time"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(dbUrl string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	poolConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	log.Println("Successfully connected to the database.")
	return pool, nil
}

func RunDatabaseMigrations(dbPool *pgxpool.Pool, migrationsPath string) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	m, err := migrate.New(
		migrationsPath,
		dbPool.Config().ConnConfig.ConnString(),
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Database migrations completed successfully.")
	return nil
}
