package models

type Repository struct {
	ID             uint `json:"id" gorm:"primaryKey"`
	Name           string
	RepositoryType []string // Ou un vrai truc plutôt qu'une clé en string
}

// Déclaration d'un ENUM, a voir si on utilise ça ou pas ?
// type Repository int
// const (
//     Generic ServerState = iota
//     Docker
//     Maven
//     NPM
// )
