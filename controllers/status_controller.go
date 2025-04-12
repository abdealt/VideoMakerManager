package controllers

import (
	"videomaker/database"
	"videomaker/models"
	"videomaker/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetStatus récupère tous les statuts avec pagination
func GetStatus(c *fiber.Ctx) error {
	// Récupération des paramètres de pagination
	pagination := utils.GetPagination(c)

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	var statuses []models.Status
	var total int64

	// Compter le nombre total d'enregistrements
	if err := db.Model(&models.Status{}).Count(&total).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors du comptage des statuts", err.Error())
	}

	// Appliquer la pagination
	query := utils.Paginate(db, pagination)
	if err := query.Find(&statuses).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération des statuts", err.Error())
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

	return utils.PaginatedResponse(c, fiber.StatusOK, "Statuts récupérés avec succès", statuses, meta)
}

// CreateStatus crée un nouveau statut
func CreateStatus(c *fiber.Ctx) error {
	var status models.Status

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Parser le corps de la requête
	if err := c.BodyParser(&status); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest,
			"Données invalides", err.Error())
	}

	// Valider les données (à implémenter si nécessaire)

	// Créer le statut dans la base de données
	if err := db.Create(&status).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la création du statut", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Statut créé avec succès", status)
}

// GetStatusByID récupère un statut par ID
func GetStatusByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var status models.Status

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	if err := db.First(&status, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound,
				"Statut non trouvé", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération du statut", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Statut trouvé", status)
}

// UpdateStatus met à jour un statut
func UpdateStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var status models.Status

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Vérifier si le statut existe
	if err := db.First(&status, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound,
				"Statut non trouvé", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération du statut", err.Error())
	}

	// Parser le corps de la requête
	if err := c.BodyParser(&status); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest,
			"Données invalides", err.Error())
	}

	// Mettre à jour le statut
	if err := db.Save(&status).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la mise à jour du statut", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Statut mis à jour avec succès", status)
}

// DeleteStatus supprime un statut
func DeleteStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Vérifier si le statut existe
	var status models.Status
	if err := db.First(&status, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound,
				"Statut non trouvé", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération du statut", err.Error())
	}

	// Supprimer le statut
	if err := db.Delete(&status).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la suppression du statut", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Statut supprimé avec succès", nil)
}
