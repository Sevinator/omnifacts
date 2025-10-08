package models

import (
	"time"

	"github.com/google/uuid"
)

type Artefact struct {
	ID           uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Name         string    `json:"name" gorm:"size:100;not null"`
	StorageKey   string    `json:"storage_key" gorm:"size:100;not null"`
	Manifest     Manifest  `json:"manifest" gorm:"foreignKey:ManifestUUID"`
	ManifestUUID uuid.UUID
	// Metadata   map[string]string

	CreatedAt time.Time // Automatically filled by gorm
	UpdatedAt time.Time // Automatically filled by gorm
}
