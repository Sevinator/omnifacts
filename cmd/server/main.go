package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"omnifacts/internal/api"
	"omnifacts/internal/config"
	"omnifacts/internal/db"
	"omnifacts/internal/service"
	"omnifacts/pkg/database"
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
	database.Connect(cfg)
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
