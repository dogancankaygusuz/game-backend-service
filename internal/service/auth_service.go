package service

import (
	"errors"
	"time"

	"github.com/dogancankaygusuz/game-backend-service/internal/domain"
	"github.com/dogancankaygusuz/game-backend-service/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("super-secret-game-key-2025") // Normalde .env'den gelmeli

// Register: Kayıt işlemini yönetir
func Register(username, password string) (*domain.Player, error) {
	// Şifreyi hashle
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	player := &domain.Player{
		ID:       uuid.New().String(),
		Username: username,
		Password: string(hashedPassword),
		Score:    0,
	}

	// Veritabanına yaz
	if err := repository.CreatePlayer(player); err != nil {
		return nil, errors.New("username already taken")
	}

	return player, nil
}

// Login: Giriş işlemini yönetir ve Token üretir
func Login(username, password string) (string, error) {
	// Kullanıcıyı bul
	player, err := repository.FindPlayerByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Şifreyi doğrula
	if err := bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Token üret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"player_id": player.ID,
		"username":  player.Username,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(jwtSecret)
}
