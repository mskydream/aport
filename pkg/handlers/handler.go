package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/aport/pkg/services"
)

type Handler struct {
	serv *services.Service
}

func NewHandler(serv *services.Service) *Handler {
	return &Handler{serv: serv}
}

func (h *Handler) InitApi(app *fiber.App) {
	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.Post("sign_up", h.signUp)
			}
		}
	}
}
