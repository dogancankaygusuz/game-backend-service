package domain

import (
	"time"

	"gorm.io/gorm"
)

// Player veritabanındaki oyuncu tablosudur
type Player struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"uniqueIndex;not null" json:"username"`
	Password  string         `json:"-"`
	Score     int            `json:"score"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
