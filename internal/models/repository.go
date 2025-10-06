package models

import (
	"github.com/google/uuid"
)

type Repository struct {
	ID            uuid.UUID `json:"id" gorm:"primaryKey"`
	Name          string
	ArtefactTypes []ArtefactType
	Type          RepositoryType
	Env           RepositoryEnv
}

type RepositoryType string
type RepositoryEnv string

const (
	Local   RepositoryType = "local"
	Remote  RepositoryType = "remote"
	Virtual RepositoryType = "virtual"
)

const (
	Dev  RepositoryEnv = "dev"
	Prod RepositoryEnv = "prod"
)
