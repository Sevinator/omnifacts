package database

import (
	"fmt"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"omnifacts/internal/config"
	"omnifacts/internal/models"
)

var DB *gorm.DB

// Connect établit la connexion à la base de données
func Connect(cfg *config.Config) {

	var err error
	slog.Info("Using variable DB_HOST with value : " + cfg.DBHost)
	slog.Info("Using variable DB_USER with value : " + cfg.DBUser)
	slog.Info("Using variable DB_PASSWORD with value : " + cfg.DBPassword)
	slog.Info("Using variable DB_NAME with value : " + cfg.DBName)
	slog.Info("Using variable DB_PORT with value : " + cfg.DBPort)

	// Configuration de la connexion
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	// Connexion à la base de données
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		slog.Error("Erreur de connexion à la base de données : " + err.Error())
	}

	slog.Info("Connexion à la base de données établie")
}

// Migrate effectue les migrations automatiques
func Migrate() {
	DB.Exec("CREATE TYPE repository_type AS ENUM ('generic', 'virtual')")

	err := DB.AutoMigrate(&models.Layer{}, &models.Manifest{}, &models.Artefact{}, &models.Repository{}) //
	if err != nil {
		slog.Error("Une erreur est survenue pendant l'application des migrations : " + err.Error())
	}
	slog.Info("Migrations effectuées")
}
