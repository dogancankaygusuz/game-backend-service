package repository

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB global değişkeni
var DB *gorm.DB

// ConnectDB veritabanı bağlantısını kurar
func ConnectDB(dbPath string) {
	var err error

	// SQLite bağlantısını aç (Pure Go sürücüsü ile)
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("🚀 Connected to SQLite database successfully! (Pure Go Mode)")
}
