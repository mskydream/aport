package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/aport/pkg/handlers"
)

func main() {
	app := fiber.New()

	err := handlers.Start(app)
	if err != nil {
		panic(err)
	}
}
