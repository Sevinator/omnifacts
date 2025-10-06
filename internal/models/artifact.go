package models

import (
	"github.com/google/uuid"
)

type Artefact struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey"`
	Name       string
	Manifest   Manifest
	Metadata   map[string]string
	StorageKey string
}

type ArtefactType string

const (
	Generic ArtefactType = "generic"
	Docker  ArtefactType = "docker"
	// Maven
	// NPM
)
