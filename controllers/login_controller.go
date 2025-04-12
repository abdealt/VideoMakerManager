// controllers/login_controller.go
package controllers

import (
	"videomaker/database"
	"videomaker/models"
	"videomaker/pkg/auth"
	"videomaker/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// LoginUserController gère l'authentification des utilisateurs
func LoginUserController(c *fiber.Ctx) error {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Récupérer les données de la requête
	if err := c.BodyParser(&loginData); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest,
			"Erreur dans les données envoyées", err.Error())
	}

	// Validation des données
	if loginData.Username == "" || loginData.Password == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest,
			"Nom d'utilisateur et mot de passe requis", "")
	}

	// Récupérer l'utilisateur de la base de données
	var user models.User
	if err := db.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized,
			"Identifiants invalides", "")
	}

	// Comparer le mot de passe fourni avec le mot de passe haché
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized,
			"Identifiants invalides", "")
	}

	// Générer un token JWT
	token, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la génération du token", err.Error())
	}

	// Retourner le token et les informations de l'utilisateur
	return utils.SuccessResponse(c, fiber.StatusOK, "Connexion réussie", fiber.Map{
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"name":     user.Name,
		},
		"token": token,
	})
}
