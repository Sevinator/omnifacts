package api

import (
	"log/slog"
	"net/http"
	"omnifactory/internal/service"
	"omnifactory/internal/utils"

	"github.com/gorilla/mux"
)

func SetupRoutes(artefactService service.ArtefactService) *mux.Router {
	router := mux.NewRouter()

	// Créer les handlers
	artefactHandler := NewArtefactHandler(artefactService)

	// Définir les routes API
	api := router.PathPrefix("/api").Subrouter()

	// Routes pour les tâches
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("On est là hein !")
	}).Methods("GET")

	api.HandleFunc("/artefact", artefactHandler.CreateArtefact).Methods("POST")
	api.HandleFunc("/artefact", artefactHandler.GetArtefacts).Methods("GET")
	api.HandleFunc("/artefact/{id}", artefactHandler.UpdateArtefact).Methods("PUT")
	api.HandleFunc("/artefact/{id}", artefactHandler.DeleteArtefact).Methods("DELETE")

	// Route de santé
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteSuccess(w, 200, map[string]string{"status": "OK"})
	}).Methods("GET")

	return router
}
