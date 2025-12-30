package handler

import (
	"github.com/dogancankaygusuz/game-backend-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(c *fiber.Ctx) error {
	var req AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	player, err := service.Register(req.Username, req.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(player)
}

func LoginHandler(c *fiber.Ctx) error {
	var req AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}

	token, err := service.Login(req.Username, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"token": token})
}
