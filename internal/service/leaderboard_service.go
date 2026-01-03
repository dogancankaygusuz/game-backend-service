package service

import (
	"github.com/dogancankaygusuz/game-backend-service/internal/domain"
	"github.com/dogancankaygusuz/game-backend-service/internal/repository"
)

// SubmitScore: Eğer yeni skor eskisinden yüksekse günceller
func SubmitScore(playerID string, score int) (*domain.Player, error) {
	// 1. Mevcut oyuncuyu bul
	player, err := repository.FindPlayerByID(playerID)
	if err != nil {
		return nil, err
	}

	// 2. Eğer yeni skor daha yüksekse güncelle
	if score > player.Score {
		err := repository.UpdateScore(playerID, score)
		if err != nil {
			return nil, err
		}
		player.Score = score // Dönüş değeri için objeyi de güncelle
	}

	return player, nil
}

// GetLeaderboard: Sıralamayı getirir
func GetLeaderboard() ([]domain.Player, error) {
	// İlk 10 oyuncuyu getir
	return repository.GetTopPlayers(10)
}
