package middleware

import (
	"strings"
	"videomaker/pkg/auth"
	"videomaker/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// Protected est un middleware qui protège les routes en vérifiant le token JWT
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Récupérer l'en-tête Authorization
		authHeader := c.Get("Authorization")

		// Vérifier si l'en-tête Authorization est présent
		if authHeader == "" {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized,
				"Accès non autorisé", "Token d'authentification manquant")
		}

		// Extraire le token du format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized,
				"Accès non autorisé", "Format de token invalide")
		}

		// Valider le token
		claims, err := auth.ValidateToken(tokenParts[1])
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized,
				"Accès non autorisé", err.Error())
		}

		// Stocker les informations de l'utilisateur dans le contexte de la requête
		c.Locals("userID", claims.UserID)
		c.Locals("username", claims.Username)

		// Continuer vers la prochaine étape
		return c.Next()
	}
}
