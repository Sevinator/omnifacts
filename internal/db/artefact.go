package db

import (
	"omnifactory/internal/models"

	"gorm.io/gorm"
)

// Interface de Gestion en base
type ArtefactDB interface {
	Create(artefact *models.Artefact) error
	Retrieve() ([]models.Artefact, error)
	Update(artefact *models.Artefact) error
	Delete(id uint) error
}

// Structure gérant la base
type artefactDB struct {
	db *gorm.DB
}

func NewArtefactDB(db *gorm.DB) ArtefactDB {
	return &artefactDB{db: db}
}

// Implémentation de l'interface
// Create
func (r *artefactDB) Create(artefact *models.Artefact) error {
	return r.db.Create(artefact).Error
}

// Retrieve
func (r *artefactDB) Retrieve() ([]models.Artefact, error) {
	var artefacts []models.Artefact
	err := r.db.Find(&artefacts).Error
	return artefacts, err
}

// Update
func (r *artefactDB) Update(artefact *models.Artefact) error {
	return r.db.Save(artefact).Error
}

// Delete
func (r *artefactDB) Delete(id uint) error {
	return r.db.Delete(&models.Artefact{}, id).Error
}
