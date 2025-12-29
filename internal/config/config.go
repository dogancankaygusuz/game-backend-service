package config

import (
	"os"
	"strconv"
)

// Config yapısı tüm uygulama ayarlarını tutar
type Config struct {
	ServerPort string
	DBPath     string // SQLite dosya yolu
}

// LoadConfig ortam değişkenlerini okur veya varsayılanları atar
func LoadConfig() *Config {
	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBPath:     getEnv("DB_PATH", "game.db"), // Proje ana dizininde oluşacak
	}
}

// Yardımcı fonksiyon: Env oku yoksa default değeri dön
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// String'i int'e çeviren yardımcı (İleride lazım olabilir)
func getEnvAsInt(key string, fallback int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return fallback
}
