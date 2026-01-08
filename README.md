# 🎮 Game Backend Service

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Framework](https://img.shields.io/badge/Fiber-v2-black?style=flat)
![Database](https://img.shields.io/badge/SQLite-GORM-blue?style=flat)
![License](https://img.shields.io/badge/License-MIT-green)

High-performance, scalable, and secure backend service designed for multiplayer games. Built with **Go**, **Fiber**, and **Clean Architecture** principles.

## 🚀 Proje Hakkında
Bu proje, mobil veya PC oyunları için gerekli olan merkezi sunucu ihtiyaçlarını karşılamak üzere geliştirilmiştir. Oyuncu kimlik doğrulama, güvenli skor takibi, liderlik tablosu ve temel hile koruma (Anti-Cheat) mekanizmalarını içerir.

"Server-Side Authoritative" (Sunucu Tabanlı Otorite) yaklaşımı benimsenerek, istemci tarafındaki manipülasyonların önüne geçilmesi hedeflenmiştir.

## ✨ Temel Özellikler

- **🔐 Kimlik Doğrulama (Auth):**
  - JWT (JSON Web Token) tabanlı güvenli oturum yönetimi.
  - Bcrypt ile şifrelerin hashlenerek saklanması.
  - Middleware ile korumalı rotalar.

- **🏆 Liderlik Tablosu (Leaderboard):**
  - Gerçek zamanlı skor güncelleme.
  - En yüksek puana sahip oyuncuların listelenmesi (Top 10).
  - Skorun sadece rekor kırıldığında güncellenmesi mantığı.

- **🛡️ Güvenlik ve Anti-Cheat:**
  - **Rate Limiting:** IP tabanlı hız sınırı ile Spam/DDoS koruması (Dakikada max 20 istek).
  - **Logic Validation:** Negatif veya imkansız skor gönderimlerini engelleyen mantıksal kontroller.

- **⚙️ Mimari ve DevOps:**
  - **Clean Architecture:** Katmanlı mimari (Handler -> Service -> Repository -> Domain).
  - **Graceful Shutdown:** Sunucu kapanırken veri kaybını önleyen güvenli kapanış mekanizması.
  - **SQLite (Pure Go):** CGO gerektirmeyen, taşınabilir veritabanı yapısı.

## 🛠️ Teknoloji Yığını

- **Dil:** Go (Golang)
- **Web Framework:** Fiber v2 (Express.js benzeri yüksek performanslı yapı)
- **Veritabanı:** SQLite (GORM ORM ile)
- **Konfigürasyon:** Standart Go yapılandırması
- **Güvenlik:** JWT, Rate Limiter

## 📂 Proje Mimarisi

game-backend-service/
├── cmd/server/ # Uygulamanın giriş noktası (Main)
├── internal/
│ ├── config/ # Ortam değişkenleri ve ayarlar
│ ├── domain/ # Veri modelleri (Structs)
│ ├── handler/ # HTTP isteklerini karşılayan katman
│ ├── middleware/ # JWT ve Rate Limit ara katmanları
│ ├── repository/ # Veritabanı işlemleri
│ └── service/ # İş mantığı (Business Logic)
└── go.mod # Bağımlılık yönetimi

## 🔌 API Dokümantasyonu

| Metot | Endpoint | Açıklama | Auth Gerekli? |
|-------|----------|----------|---------------|
| `POST` | `/auth/register` | Yeni oyuncu kaydı oluşturur | ❌ Hayır |
| `POST` | `/auth/login` | Giriş yapar ve Token döner | ❌ Hayır |
| `GET` | `/health` | Sunucu sağlık durumunu kontrol eder | ❌ Hayır |
| `GET` | `/api/profile` | Oyuncunun kendi profilini getirir | ✅ Evet (Token) |
| `POST` | `/api/leaderboard/submit` | Yeni skor gönderir | ✅ Evet (Token) |
| `GET` | `/api/leaderboard/top` | En iyi 10 oyuncuyu listeler | ✅ Evet (Token) |

## 🚀 Kurulum ve Çalıştırma

### Gereksinimler
- Go 1.18 veya üzeri

### Adımlar

1. **Projeyi Klonlayın:**
   ```bash
   git clone https://github.com/dogancankaygusuz/game-backend-service.git
   cd game-backend-service
   ```

2. **Bağımlılıkları Yükleyin:**
    ```bash
    go mod tidy
    ```

3. **Sunucuyu Başlatın:**
    ```bash
    go run cmd/server/main.go
    ```

4. **Test Edin::**
    Sunucu http://localhost:8080 adresinde çalışacaktır. Postman veya cURL kullanarak istek atabilirsiniz.