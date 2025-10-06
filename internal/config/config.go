package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() *Config {
	// Charger le fichier .env
	err := godotenv.Load()
	if err != nil {
		slog.Error("Fichier .env non trouvé, utilisation des variables d'environnement")
	}

	return &Config{
		Port: getEnv("PORT", "6666"),
		// DBHost:     getEnv("DB_HOST", "localhost"),
		// DBPort:     getEnv("DB_PORT", "5432"),
		// DBUser:     getEnv("DB_USER", "postgres"),
		// DBPassword: getEnv("DB_PASSWORD", ""),
		// DBName:     getEnv("DB_NAME", "todo_api"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
