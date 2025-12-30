package handler

import (
	"github.com/dogancankaygusuz/game-backend-service/internal/repository"
	"github.com/gofiber/fiber/v2"
)

// GetProfile -> GET /player/profile
func GetProfile(c *fiber.Ctx) error {
	// Middleware'den gelen player_id'yi al
	playerID := c.Locals("player_id").(string)

	// Veritabanından oyuncuyu bul
	player, err := repository.FindPlayerByID(playerID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Player not found"})
	}

	// Oyuncuyu döndür (Şifre zaten JSON'da gizli - domain modelinde ayarlamıştık)
	return c.Status(200).JSON(player)
}
