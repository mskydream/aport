package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/aport/config"
	"github.com/mskydream/aport/db"
	"github.com/mskydream/aport/pkg/handlers"
	"github.com/mskydream/aport/pkg/repositories"
	"github.com/mskydream/aport/pkg/services"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("config error")
	}

	db, err := db.PostgreSQLConnection(&cfg)
	if err != nil {
		panic(err)
	}

	// fmt.Println(db)

	repos := repositories.NewRepository(db)
	services := services.NewService(repos)
	handlers := handlers.NewHandler(services)

	app := fiber.New()
	// app.Use(logger.New())
	handlers.RegisterHandlers(app)
	log.Fatal(app.Listen(cfg.Port))
}
