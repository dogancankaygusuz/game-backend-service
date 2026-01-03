package handler

import (
	"github.com/dogancankaygusuz/game-backend-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ScoreRequest struct {
	Score int `json:"score"`
}

// SubmitScoreHandler -> POST /api/leaderboard/submit
func SubmitScoreHandler(c *fiber.Ctx) error {
	// Middleware sayesinde ID'yi biliyoruz
	playerID := c.Locals("player_id").(string)

	var req ScoreRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid score format"})
	}

	// Servisi çağır
	player, err := service.SubmitScore(playerID, req.Score)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update score"})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":            "Score processed",
		"current_high_score": player.Score,
	})
}

// GetLeaderboardHandler -> GET /api/leaderboard/top
func GetLeaderboardHandler(c *fiber.Ctx) error {
	players, err := service.GetLeaderboard()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch leaderboard"})
	}

	return c.Status(200).JSON(players)
}
