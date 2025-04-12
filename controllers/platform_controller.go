package controllers

import (
	"videomaker/database"
	"videomaker/models"
	"videomaker/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetPlatforms récupère toutes les plateformes avec pagination
func GetPlatforms(c *fiber.Ctx) error {
	// Récupération des paramètres de pagination
	pagination := utils.GetPagination(c)

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	var platforms []models.Platform
	var total int64

	// Compter le nombre total d'enregistrements
	if err := db.Model(&models.Platform{}).Count(&total).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors du comptage des plateformes", err.Error())
	}

	// Appliquer la pagination
	query := utils.Paginate(db, pagination)
	if err := query.Find(&platforms).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération des plateformes", err.Error())
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

	return utils.PaginatedResponse(c, fiber.StatusOK, "Plateformes récupérées avec succès", platforms, meta)
}

// CreatePlatform crée une nouvelle plateforme
func CreatePlatform(c *fiber.Ctx) error {
	var platform models.Platform

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Parser le corps de la requête
	if err := c.BodyParser(&platform); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest,
			"Données invalides", err.Error())
	}

	// Valider les données (à implémenter si nécessaire)

	// Créer la plateforme dans la base de données
	if err := db.Create(&platform).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la création de la plateforme", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Plateforme créée avec succès", platform)
}

// GetPlatformByID récupère une plateforme par ID
func GetPlatformByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var platform models.Platform

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	if err := db.First(&platform, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound,
				"Plateforme non trouvée", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération de la plateforme", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Plateforme trouvée", platform)
}

// UpdatePlatform met à jour une plateforme
func UpdatePlatform(c *fiber.Ctx) error {
	id := c.Params("id")
	var platform models.Platform

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Vérifier si la plateforme existe
	if err := db.First(&platform, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound,
				"Plateforme non trouvée", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération de la plateforme", err.Error())
	}

	// Parser le corps de la requête
	if err := c.BodyParser(&platform); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest,
			"Données invalides", err.Error())
	}

	// Mettre à jour la plateforme
	if err := db.Save(&platform).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la mise à jour de la plateforme", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Plateforme mise à jour avec succès", platform)
}

// DeletePlatform supprime une plateforme
func DeletePlatform(c *fiber.Ctx) error {
	id := c.Params("id")

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Vérifier si la plateforme existe
	var platform models.Platform
	if err := db.First(&platform, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound,
				"Plateforme non trouvée", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération de la plateforme", err.Error())
	}

	// Supprimer la plateforme
	if err := db.Delete(&platform).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la suppression de la plateforme", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Plateforme supprimée avec succès", nil)
}
