package models

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

type RepositoryType string

const (
	Generic RepositoryType = "generic"
	Virtual RepositoryType = "virtual"
)

// Implémentation de la méthode String() pour afficher la valeur en tant que chaîne de caractères
func (s RepositoryType) String() string {
	return string(s)
}

// Implémentation de la méthode Value() pour GORM
// Cela permet à GORM de stocker la valeur dans la base de données en tant que type approprié
func (s RepositoryType) Value() (driver.Value, error) {
	return string(s), nil
}

// Implémentation de la méthode Scan() pour GORM
// Cela permet de récupérer la valeur depuis la base de données sous forme d'énum
func (s *RepositoryType) Scan(value interface{}) error {
	*s = RepositoryType(value.(string))
	return nil
}

type Repository struct {
	ID              uuid.UUID        `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Name            string           `json:"name" gorm:"size:100;not null"`
	RepositoryTypes []RepositoryType `json:"repository_type" gorm:"type:repository_type;default:'generic'"`
}

// Déclaration d'un ENUM, a voir si on utilise ça ou pas ?
// type Repository int
// const (
//     Generic ServerState = iota
//     Docker
//     Maven
//     NPM
// )
