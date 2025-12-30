package main

import (
	"log"

	"github.com/dogancankaygusuz/game-backend-service/internal/config"
	"github.com/dogancankaygusuz/game-backend-service/internal/domain"
	"github.com/dogancankaygusuz/game-backend-service/internal/handler"
	"github.com/dogancankaygusuz/game-backend-service/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.LoadConfig()
	repository.ConnectDB(cfg.DBPath)

	// Veritabanı Tablosunu Oluştur (Migration)
	repository.DB.AutoMigrate(&domain.Player{})
	log.Println("✅ Database migrations completed")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	// Rotalar
	auth := app.Group("/auth")
	auth.Post("/register", handler.RegisterHandler)
	auth.Post("/login", handler.LoginHandler)

	log.Fatal(app.Listen(":" + cfg.ServerPort))
}
