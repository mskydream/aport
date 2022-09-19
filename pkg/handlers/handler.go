package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/aport/models"
	"github.com/mskydream/aport/pkg/services"
)

type Handler struct {
	serv *services.Service
}

func NewHandler(serv *services.Service) *Handler {
	return &Handler{serv: serv}

}

func (h *Handler) RegisterHandlers(app *fiber.App) {
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

func (h *Handler) signUp(ctx *fiber.Ctx) error {
	var input models.UserProfile

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.ResponseData{Status: false, Message: "incorrect input", Data: err})
	}

	id, err := h.serv.SignUp(&input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.ResponseData{Status: false, Message: "server error", Data: err})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.ResponseData{Status: true, Message: "sucess", Data: id})
}
