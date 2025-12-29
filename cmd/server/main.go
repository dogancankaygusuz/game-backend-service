package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Fiber instance oluştur
	app := fiber.New(fiber.Config{
		AppName: "Game Backend Service",
	})

	// Middleware'ler
	app.Use(logger.New())  // İstekleri loglar
	app.Use(recover.New()) // Panic durumunda sunucunun çökmesini önler

	// Basit bir Health Check rotası
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "online",
			"message": "Game Backend Service is running correctly 🚀",
		})
	})

	// Sunucuyu başlat
	log.Fatal(app.Listen(":8080"))
}
