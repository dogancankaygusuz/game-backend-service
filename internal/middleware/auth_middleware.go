package middleware

import (
	"strings"

	"github.com/dogancankaygusuz/game-backend-service/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-game-key-2025") // Service ile aynı key olmalı!

// Protected: Sadece geçerli token'ı olanların geçmesine izin verir
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Config'i yükle (Gizli anahtarı almak için)
		cfg := config.LoadConfig()

		// 2. Header'dan token'ı al
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized: Missing token"})
		}

		// "Bearer <token>" formatını temizle
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// 3. Token'ı doğrula
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// İmza yöntemini kontrol et (HMAC olmalı)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			// DÜZELTME BURADA: Config'den gelen anahtarı kullanıyoruz
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized: Invalid token"})
		}

		// 4. Token içindeki verileri (Claims) al
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized: Invalid claims"})
		}

		// 5. User ID'yi context'e kaydet
		c.Locals("player_id", claims["player_id"])

		// Yola devam et
		return c.Next()
	}
}
