package main

import (
	"database/sql"
	"embed"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"todo-list/backend/repository"

	_ "github.com/lib/pq"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	db, dbErr := sql.Open("postgres", "postgresql://todo-user:password@localhost:5434/todo-db?sslmode=disable")

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	defer func(db *sql.DB) {
		dbErr := db.Close()
		if dbErr != nil {
			log.Fatal(dbErr)
		}
	}(db)

	driver, dbErr := postgres.WithInstance(db, &postgres.Config{})
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	m, dbErr := migrate.NewWithDatabaseInstance(
		"file://backend/db/migrations",
		"postgres", driver)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	if dbErr := m.Up(); dbErr != nil && !errors.Is(dbErr, migrate.ErrNoChange) {
		log.Fatal(dbErr)
	}

	taskRepo := repository.NewTaskRepository(db)

	app := NewApp(taskRepo)

	err := wails.Run(&options.App{
		Title:  "todo-list",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
