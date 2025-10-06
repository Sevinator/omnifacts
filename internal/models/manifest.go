package models

import (
	"time"

	"github.com/google/uuid"
)

// A redéfinir par rapport à la norme
type Manifest struct {
	ID           uuid.UUID
	OwnerID      uuid.UUID
	LastDownload time.Time
	UploadedDate time.Time
	LastUpdated  time.Time
	Layers       []string
	Author       string
}
