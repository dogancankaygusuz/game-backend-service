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
