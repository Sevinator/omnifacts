package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"omnifactory/internal/models"
	"omnifactory/internal/service"
	"omnifactory/internal/utils"

	"github.com/gorilla/mux"
)

type ArtefactHandler struct {
	artefactService service.ArtefactService
}

func NewArtefactHandler(artefactService service.ArtefactService) *ArtefactHandler {
	return &ArtefactHandler{artefactService: artefactService}
}

// POST
func (h *ArtefactHandler) CreateArtefact(w http.ResponseWriter, r *http.Request) {
	var artefact models.Artefact

	// Décoder le JSON de la requête
	if err := json.NewDecoder(r.Body).Decode(&artefact); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Format JSON invalide")
		return
	}

	// Créer la tâche
	if err := h.artefactService.CreateArtefact(&artefact); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusCreated, artefact)
}

// GET
func (h *ArtefactHandler) GetArtefacts(w http.ResponseWriter, r *http.Request) {
	artefacts, err := h.artefactService.RetrieveArtefacts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, artefacts)
}

// Put
func (h *ArtefactHandler) UpdateArtefact(w http.ResponseWriter, r *http.Request) {
	var artefact models.Artefact
	if err := json.NewDecoder(r.Body).Decode(&artefact); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Format JSON invalide")
		return
	}

	if err := h.artefactService.UpdateArtefact(&artefact); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, map[string]string{"message": "Artefact mis à jour"})
}

// DeleteArtefact supprime une tâche
func (h *ArtefactHandler) DeleteArtefact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "ID invalide")
		return
	}

	if err := h.artefactService.DeleteArtefact(uint(id)); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteSuccess(w, http.StatusOK, map[string]string{"message": "Tâche supprimée"})
}
