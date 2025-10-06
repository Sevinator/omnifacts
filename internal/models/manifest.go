package models

import (
	"time"
)

// A redéfinir par rapport à la norme
type Manifest struct {
	ID           uint
	OwnerID      uint
	LastDownload time.Time
	UploadedDate time.Time
	LastUpdated  time.Time
	Layers       []string
	Author       string
}
