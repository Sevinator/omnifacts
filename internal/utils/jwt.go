package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(getJWTSecret())

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// En production, ceci devrait générer une erreur !
		return "ma-cle-secrete-super-sure"
	}
	return secret
}

// GenerateToken génère un token JWT pour un utilisateur
func GenerateToken(userID uint, email string) (string, error) {
	// Définir les claims (revendications)
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24h
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "todo-api",
		},
	}

	// Créer le token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signer le token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken valide un token JWT et retourne les claims
func ValidateToken(tokenString string) (*Claims, error) {
	// Parser le token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Vérifier la méthode de signature
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("méthode de signature inattendue")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Vérifier si le token est valide
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token invalide")
}
