package models

import (
	"time"

	"github.com/google/uuid"
)

// A redéfinir par rapport à la norme
type Manifest struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Author     *string   `json:"author" gorm:"size 100;"`
	Layers     []Layer   `json:"layers" gorm:"many2many:manifest_layers;"` // gorm:"foreignKey:LayerRefer"
	ArtefactID uint      `json:"artefact"`
	// Gorm auto filled
	LastDownload time.Time
	LastUpdated  time.Time
	UploadedDate time.Time
	// OwnerID uint    `json:"owner_id" gorm:""`
}
