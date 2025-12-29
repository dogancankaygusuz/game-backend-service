package main

import (
	"log"

	"github.com/dogancankaygusuz/game-backend-service/internal/config"
	"github.com/dogancankaygusuz/game-backend-service/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// 1. Konfigürasyonu Yükle
	cfg := config.LoadConfig()

	// 2. Veritabanına Bağlan
	repository.ConnectDB(cfg.DBPath)

	// 3. Fiber Uygulamasını Başlat
	app := fiber.New(fiber.Config{
		AppName: "Game Backend Service",
	})

	// Middleware'ler
	app.Use(logger.New())
	app.Use(recover.New())

	// Rotalar
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "online",
			"message": "Game Backend Service is running",
			"db":      "connected (sqlite)",
		})
	})

	// Sunucuyu Dinle
	log.Printf("Server starting on port %s", cfg.ServerPort)
	log.Fatal(app.Listen(":" + cfg.ServerPort))
}
