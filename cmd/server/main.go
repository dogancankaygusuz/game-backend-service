package main

import (
	"log"
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

	// 2. Uygulamayı Başlat
	app := fiber.New()

	// 3. Temel Middleware'ler
	app.Use(logger.New())  // Log tutar
	app.Use(recover.New()) // Çökme önleyici

	// --- RATE LIMITER (Hız Sınırı - YENİ EKLENEN KISIM) ---
	// Her IP adresi için dakikada maksimum 20 istek
	app.Use(limiter.New(limiter.Config{
		Max:        20,              // 1 dakikada en fazla 20 istek
		Expiration: 1 * time.Minute, // Engelleme süresi
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Engellemeyi IP adresine göre yap
		},
		LimitReached: func(c *fiber.Ctx) error {
			// Limit aşılırsa bu hatayı döndür
			return c.Status(429).JSON(fiber.Map{
				"error": "Too many requests. Please slow down.",
			})
		},
	}))
	// -------------------------------------------------------

	// 4. Rotalar (Routes)

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

	api.Get("/profile", handler.GetProfile)
	api.Post("/leaderboard/submit", handler.SubmitScoreHandler)
	api.Get("/leaderboard/top", handler.GetLeaderboardHandler)

	// 5. Sunucuyu Dinle
	log.Printf("Server starting on port %s", cfg.ServerPort)
	log.Fatal(app.Listen(":" + cfg.ServerPort))
}
