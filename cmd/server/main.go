package main

import (
	"log"

	"github.com/dogancankaygusuz/game-backend-service/internal/config"
	"github.com/dogancankaygusuz/game-backend-service/internal/domain"
	"github.com/dogancankaygusuz/game-backend-service/internal/handler"
	"github.com/dogancankaygusuz/game-backend-service/internal/middleware"
	"github.com/dogancankaygusuz/game-backend-service/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.LoadConfig()
	repository.ConnectDB(cfg.DBPath)

	repository.DB.AutoMigrate(&domain.Player{})
	log.Println("✅ Database migrations completed")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	// --- PUBLIC ROUTES (Herkes erişebilir) ---
	auth := app.Group("/auth")
	auth.Post("/register", handler.RegisterHandler)
	auth.Post("/login", handler.LoginHandler)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"status": "online"})
	})

	// --- PROTECTED ROUTES (Sadece Token ile erişilebilir) ---
	// "middleware.Protected()" bu grubun bekçisidir
	api := app.Group("/api", middleware.Protected())

	api.Get("/profile", handler.GetProfile) // <-- YENİ ROTA
	api.Post("/leaderboard/submit", handler.SubmitScoreHandler)
	api.Get("/leaderboard/top", handler.GetLeaderboardHandler)
	log.Fatal(app.Listen(":" + cfg.ServerPort))
}
