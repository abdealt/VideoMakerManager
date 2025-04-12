package controllers

import (
	"videomaker/database"
	"videomaker/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Vérification du mot de passe lors de la connexion
func LoginUserController(c *fiber.Ctx) error {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	// Récupérer les données de la requête
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erreur dans les données envoyées")
	}

	// Récupérer l'utilisateur de la base de données
	var user models.User
	if err := db.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Utilisateur non trouvé")
	}

	// Comparer le mot de passe fourni avec le mot de passe haché
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Mot de passe incorrect")
	}

	// Si le mot de passe est correct, tu peux générer un token ou effectuer d'autres actions
	return c.SendString("Connexion réussie")
}
