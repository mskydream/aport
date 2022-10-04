package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mskydream/aport/config"
	"github.com/mskydream/aport/db"
	"github.com/mskydream/aport/pkg/repositories"
	"github.com/mskydream/aport/pkg/services"
)

func Start(app *fiber.App) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	db, err := db.PostgreSQLConnection(&cfg)
	if err != nil {
		return err
	}

	repos := repositories.NewRepository(db)
	services := services.NewService(repos)
	handlers := NewHandler(services)

	app.Use(logger.New())
	handlers.InitApi(app)

	log.Fatal(app.Listen(cfg.Port))
	err = app.Listen(cfg.Port)

	return err
}
