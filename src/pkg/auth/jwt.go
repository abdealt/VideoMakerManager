package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// JWTClaims représente les claims JWT
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Variables d'environnement et constantes
var (
	jwtSecret     []byte
	tokenDuration = 24 * time.Hour // Durée de validité du token
)

// Init initialise le module JWT
func Init() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("erreur lors du chargement du fichier .env: %v", err)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET")
	if jwtSecretKey == "" {
		// Utiliser une valeur par défaut uniquement en développement
		jwtSecretKey = "default_jwt_secret_key_change_in_production"
	}

	jwtSecret = []byte(jwtSecretKey)
	return nil
}

// GenerateToken génère un token JWT pour un utilisateur
func GenerateToken(userID uint, username string) (string, error) {
	// Créer les claims
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// Créer le token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signer le token
	return token.SignedString(jwtSecret)
}

// ValidateToken valide un token JWT
func ValidateToken(tokenString string) (*JWTClaims, error) {
	// Parser le token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Vérifier la méthode de signature
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("méthode de signature inattendue: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	// Extraire les claims
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("token invalide")
}
