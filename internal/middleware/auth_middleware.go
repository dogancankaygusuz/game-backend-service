package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-game-key-2025") // Service ile aynı key olmalı!

// Protected: Sadece geçerli token'ı olanların geçmesine izin verir
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Header'dan token'ı al
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized: Missing token"})
		}

		// "Bearer <token>" formatını temizle
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// 2. Token'ı doğrula
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized: Invalid token"})
		}

		// 3. Token içindeki verileri (Claims) al
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized: Invalid claims"})
		}

		// 4. User ID'yi context'e (Locals) kaydet ki sonraki handler kullanabilsin
		c.Locals("player_id", claims["player_id"])

		// Yola devam et
		return c.Next()
	}
}
