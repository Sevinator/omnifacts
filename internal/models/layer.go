package models

import "github.com/google/uuid"

type Layer struct {
	ID   uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Name string    `json:"name" gorm:"size:100;not null"`
	Hash string    `json:"hash" gorm:"size:100;not null"`
}
