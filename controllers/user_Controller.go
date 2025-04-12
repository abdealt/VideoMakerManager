package controllers

import (
	"videomaker/database"
	"videomaker/models"
	"videomaker/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GetUsers récupère tous les utilisateurs avec pagination
func GetUsers(c *fiber.Ctx) error {
	// Récupération des paramètres de pagination
	pagination := utils.GetPagination(c)

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	var users []models.User
	var total int64

	// Compter le nombre total d'enregistrements
	if err := db.Model(&models.User{}).Count(&total).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors du comptage des utilisateurs", err.Error())
	}

	// Appliquer la pagination
	query := utils.Paginate(db, pagination)
	if err := query.Find(&users).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération des utilisateurs", err.Error())
	}

	// Calculer le nombre total de pages
	lastPage := int(total) / pagination.Limit
	if int(total)%pagination.Limit != 0 {
		lastPage++
	}

	// Créer les métadonnées de pagination
	meta := utils.Meta{
		Total:    total,
		Page:     pagination.Page,
		PageSize: pagination.Limit,
		LastPage: lastPage,
	}

	return utils.PaginatedResponse(c, fiber.StatusOK, "Utilisateurs récupérés avec succès", users, meta)
}

// CreateUser crée un nouvel utilisateur
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Parser le corps de la requête
	if err := c.BodyParser(&user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest,
			"Données invalides", err.Error())
	}

	// Valider les données (à implémenter)

	// Hacher le mot de passe avec bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors du hachage du mot de passe", err.Error())
	}

	// Remplacer le mot de passe par le mot de passe haché
	user.Password = string(hashedPassword)

	// Créer l'utilisateur dans la base de données
	if err := db.Create(&user).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la création de l'utilisateur", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Utilisateur créé avec succès", user)
}

// GetUserByID récupère un utilisateur par ID
func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound,
				"Utilisateur non trouvé", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération de l'utilisateur", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Utilisateur trouvé", user)
}

// UpdateUser met à jour un utilisateur
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Vérifier si l'utilisateur existe
	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound,
				"Utilisateur non trouvé", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération de l'utilisateur", err.Error())
	}

	// Sauvegarder le mot de passe actuel
	currentPassword := user.Password

	// Parser le corps de la requête
	if err := c.BodyParser(&user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest,
			"Données invalides", err.Error())
	}

	// Si le mot de passe a été modifié, le hacher
	if user.Password != "" && user.Password != currentPassword {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError,
				"Erreur lors du hachage du mot de passe", err.Error())
		}
		user.Password = string(hashedPassword)
	} else {
		// Sinon, conserver l'ancien mot de passe
		user.Password = currentPassword
	}

	// Mettre à jour l'utilisateur
	if err := db.Save(&user).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la mise à jour de l'utilisateur", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Utilisateur mis à jour avec succès", user)
}

// DeleteUser supprime un utilisateur
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Vérifier si l'utilisateur existe
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound,
				"Utilisateur non trouvé", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération de l'utilisateur", err.Error())
	}

	// Supprimer l'utilisateur
	if err := db.Delete(&user).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la suppression de l'utilisateur", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Utilisateur supprimé avec succès", nil)
}
