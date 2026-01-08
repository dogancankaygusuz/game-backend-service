package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dogancankaygusuz/game-backend-service/internal/config"
	"github.com/dogancankaygusuz/game-backend-service/internal/domain"
	"github.com/dogancankaygusuz/game-backend-service/internal/handler"
	"github.com/dogancankaygusuz/game-backend-service/internal/middleware"
	"github.com/dogancankaygusuz/game-backend-service/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// 1. Konfigürasyon ve DB Bağlantısı
	cfg := config.LoadConfig()
	repository.ConnectDB(cfg.DBPath)

	repository.DB.AutoMigrate(&domain.Player{})
	log.Println("✅ Database migrations completed")

	// 2. Uygulamayı Hazırla
	app := fiber.New(fiber.Config{
		AppName: "Game Backend Service",
	})

	// 3. Middleware'ler
	app.Use(logger.New())
	app.Use(recover.New())

	// Rate Limiter
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(fiber.Map{"error": "Too many requests"})
		},
	}))

	// 4. Rotalar
	auth := app.Group("/auth")
	auth.Post("/register", handler.RegisterHandler)
	auth.Post("/login", handler.LoginHandler)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"status": "online"})
	})

	api := app.Group("/api", middleware.Protected())
	api.Get("/profile", handler.GetProfile)
	api.Post("/leaderboard/submit", handler.SubmitScoreHandler)
	api.Get("/leaderboard/top", handler.GetLeaderboardHandler)

	// --- 5. GRACEFUL SHUTDOWN (Zarif Kapanış) MEKANİZMASI ---

	// Sunucuyu ayrı bir Goroutine'de (thread) başlatıyoruz ki ana akış bloklanmasın
	go func() {
		if err := app.Listen(":" + cfg.ServerPort); err != nil {
			log.Panic(err)
		}
	}()

	log.Printf("🚀 Server is running on port %s", cfg.ServerPort)

	// İşletim sisteminden gelecek kapatma sinyallerini dinle (Ctrl+C gibi)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Sinyal gelene kadar burada bekle (Blokla)
	<-c

	log.Println("⚠️  Shutdown signal received. Closing connection...")

	// Sunucuyu güvenli kapat
	if err := app.Shutdown(); err != nil {
		log.Printf("❌ Error shutting down server: %v", err)
	}

	// SQL Bağlantısını kapat
	sqlDB, err := repository.DB.DB()
	if err == nil {
		sqlDB.Close()
		log.Println("🔌 Database connection closed")
	}

	log.Println("👋 Server exited successfully")
}
