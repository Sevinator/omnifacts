package models

type Artefact struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	Name       string
	Manifest   Manifest
	Metadata   map[string]string
	StorageKey string
}
