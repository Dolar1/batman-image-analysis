package main

import (
	"batman/app"
	batman_db "batman/db"
	"batman/server"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	configMigrationsPath = "file://./migrations"
)

func main() {
	clientApp := cli.NewApp()
	clientApp.Name = "Batman Image Analysis Platform Lite"
	clientApp.Version = "1.0.0"

	clientApp.Commands = []*cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "start server",
			Action: func(c *cli.Context) error {
				dependency, _ := app.Init()
				defer app.StopApp()

				server.StartAPIServer(dependency)
				return nil
			},
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "run database migrations",
			Action: func(c *cli.Context) error {
				dependency, _ := app.Init()
				defer app.StopApp()
				// Run database migrations
				if err := batman_db.RunDatabaseMigrations(dependency.DBPool, configMigrationsPath); err != nil {
					log.Fatalf("Error running migrations: %v", err)
				}

				return nil
			},
		},
		{
			Name:    "rollback",
			Aliases: []string{"r"},
			Usage:   "run database migrations rollback",
			Action: func(c *cli.Context) error {
				// // Load configuration
				// cfg, err := config.Load()
				// if err != nil {
				// 	return errors.Wrap(err, "Failed to load config")
				// }

				// if err := db.RollbackLatestMigration(&db.MigrationConfig{
				// 	Driver: "postgres",
				// 	URL:    cfg.DBUrl,
				// 	Path:   configMigrationsPath,
				// }); err != nil {
				// 	return errors.Wrap(err, "rollback failed for config DB")
				// }

				return nil
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
