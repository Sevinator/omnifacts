package database

import (
	"log/slog"

	// "gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"omnifacts/internal/models"
)

var DB *gorm.DB

// Connect établit la connexion à la base de données
func Connect() {
	var err error

	// Configuration de la connexion
	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_PORT"),
	// )

	// Connexion à la base de données
	DB, err = gorm.Open(sqlite.Open("./omni.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		slog.Error("Erreur de connexion à la base de données:", err)
	}

	slog.Info("Connexion à la base de données établie")
}

// Migrate effectue les migrations automatiques
func Migrate() {
	DB.AutoMigrate(&models.Artefact{}, &models.Manifest{}, &models.Repository{})
	slog.Info("Migrations effectuées")
}
