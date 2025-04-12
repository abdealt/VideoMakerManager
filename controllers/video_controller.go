package controllers

import (
	"videomaker/database"
	"videomaker/models"
	"videomaker/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetVideos récupère toutes les vidéos avec pagination
func GetVideos(c *fiber.Ctx) error {
	// Récupération des paramètres de pagination
	pagination := utils.GetPagination(c)

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	var videos []models.Video
	var total int64

	// Compter le nombre total d'enregistrements
	if err := db.Model(&models.Video{}).Count(&total).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors du comptage des vidéos", err.Error())
	}

	// Appliquer la pagination
	query := utils.Paginate(db, pagination)
	if err := query.Find(&videos).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération des vidéos", err.Error())
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

	return utils.PaginatedResponse(c, fiber.StatusOK, "Vidéos récupérées avec succès", videos, meta)
}

// CreateVideo crée une nouvelle vidéo
func CreateVideo(c *fiber.Ctx) error {
	var video models.Video

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Parser le corps de la requête
	if err := c.BodyParser(&video); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest,
			"Données invalides", err.Error())
	}

	// Valider les données (à implémenter si nécessaire)

	// Créer la vidéo dans la base de données
	if err := db.Create(&video).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la création de la vidéo", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Vidéo créée avec succès", video)
}

// GetVideoByID récupère une vidéo par ID
func GetVideoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var video models.Video

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Préchargement des relations (platform, status, creator)
	if err := db.Preload("Platform").Preload("Status").Preload("Creator").First(&video, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound,
				"Vidéo non trouvée", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération de la vidéo", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Vidéo trouvée", video)
}

// UpdateVideo met à jour une vidéo
func UpdateVideo(c *fiber.Ctx) error {
	id := c.Params("id")
	var video models.Video

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Vérifier si la vidéo existe
	if err := db.First(&video, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound,
				"Vidéo non trouvée", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération de la vidéo", err.Error())
	}

	// Parser le corps de la requête
	if err := c.BodyParser(&video); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest,
			"Données invalides", err.Error())
	}

	// Mettre à jour la vidéo
	if err := db.Save(&video).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la mise à jour de la vidéo", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Vidéo mise à jour avec succès", video)
}

// DeleteVideo supprime une vidéo
func DeleteVideo(c *fiber.Ctx) error {
	id := c.Params("id")

	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur de connexion à la base de données", err.Error())
	}

	// Vérifier si la vidéo existe
	var video models.Video
	if err := db.First(&video, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(c, fiber.StatusNotFound,
				"Vidéo non trouvée", "")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la récupération de la vidéo", err.Error())
	}

	// Supprimer la vidéo
	if err := db.Delete(&video).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError,
			"Erreur lors de la suppression de la vidéo", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Vidéo supprimée avec succès", nil)
}
