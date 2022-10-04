package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mskydream/aport/models"
)

func (h *Handler) signUp(ctx *fiber.Ctx) error {
	var input models.UserProfile

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(400).JSON(models.ResponseData{Status: false, Message: "Incorrect input", Data: err.Error()})
	}

	id, err := h.serv.SignUp(&input)
	if err != nil {
		return ctx.Status(500).JSON(models.ResponseData{Status: false, Message: "Server error", Data: err.Error()})
	}

	return ctx.Status(200).JSON(models.ResponseData{Status: true, Message: "Success", Data: id})
}
