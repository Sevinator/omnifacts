package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"omnifactory/internal/api"
	"omnifactory/internal/config"
	"omnifactory/internal/db"
	"omnifactory/internal/service"
	"omnifactory/pkg/database"
)

// import (
// 	"log/slog"
// 	"os"
// )

// func main() {
// 	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))

// 	slog.Info("Hello, on lance le projet ....")

// }

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))

	// Charger la configuration
	cfg := config.Load()

	// Établir la connexion à la base de données
	database.Connect()
	database.Migrate()

	// Initialiser les couches
	artefactDB := db.NewArtefactDB(database.DB)
	artefactService := service.NewArtefactService(artefactDB)

	// Configurer les routes
	router := api.SetupRoutes(artefactService)

	// Configurer le serveur
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	log.Printf("Serveur démarré sur le port %s", cfg.Port)
	log.Fatal(server.ListenAndServe())
}
