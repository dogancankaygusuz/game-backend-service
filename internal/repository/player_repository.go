package repository

import (
	"github.com/dogancankaygusuz/game-backend-service/internal/domain"
)

// CreatePlayer: Yeni oyuncuyu kaydeder
func CreatePlayer(player *domain.Player) error {
	return DB.Create(player).Error
}

// FindPlayerByUsername: İsme göre oyuncu bulur
func FindPlayerByUsername(username string) (*domain.Player, error) {
	var player domain.Player
	err := DB.Where("username = ?", username).First(&player).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func FindPlayerByID(id string) (*domain.Player, error) {
	var player domain.Player
	err := DB.Where("id = ?", id).First(&player).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

// UpdateScore: Oyuncunun skorunu günceller
func UpdateScore(playerID string, newScore int) error {
	// Model ile güncelleme yaparak updated_at alanının da değişmesini sağlarız
	return DB.Model(&domain.Player{}).Where("id = ?", playerID).Update("score", newScore).Error
}

// GetTopPlayers: En yüksek puana sahip 'limit' kadar oyuncuyu getirir
func GetTopPlayers(limit int) ([]domain.Player, error) {
	var players []domain.Player
	// Puanı yüksekten düşüğe (desc) sırala ve limiti uygula
	err := DB.Order("score desc").Limit(limit).Find(&players).Error
	return players, err
}
