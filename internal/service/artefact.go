package service

import (
	"omnifacts/internal/db"
	"omnifacts/internal/models"
)

type ArtefactService interface {
	CreateArtefact(artefact *models.Artefact) error
	RetrieveArtefacts() ([]models.Artefact, error)
	UpdateArtefact(artefact *models.Artefact) error
	DeleteArtefact(id uint) error
}

type artefactService struct {
	artefactDB db.ArtefactDB
}

func NewArtefactService(artefactDB db.ArtefactDB) ArtefactService {
	return &artefactService{artefactDB: artefactDB}
}

// Implementation de l'interface
// Logique de création
func (s *artefactService) CreateArtefact(artefact *models.Artefact) error {
	return s.artefactDB.Create(artefact)
}

// Logique de retrieve
func (s *artefactService) RetrieveArtefacts() ([]models.Artefact, error) {
	artefacts, err := s.artefactDB.Retrieve()
	if err != nil {
		return nil, err
	}

	return artefacts, nil
}

// Logique métier pour un Update
func (s *artefactService) UpdateArtefact(artefact *models.Artefact) error {
	return s.artefactDB.Update(artefact)
}

// Logique métier pour un Delete
func (s *artefactService) DeleteArtefact(id uint) error {
	return s.artefactDB.Delete(id)
}
