package controllers

import (
	"github.com/abdealt/videomaker/database"
	"github.com/abdealt/videomaker/models"

	"github.com/gofiber/fiber/v2"
)

// Créer un nouvel utilisateur
func CreateVideo(c *fiber.Ctx) error {
	var video models.Video
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	// Enregistrer l'utilisateur dans la base de données avec Gorm
	if err := db.Create(&video).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la création de l'utilisateur")
	}

	return c.Status(fiber.StatusCreated).JSON(video)
}

func GetVideos(c *fiber.Ctx) error {
	// Récupère l'instance de la base de données depuis le contexte ou ailleurs
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).SendString("Erreur de connexion à la base de données")
	}

	var videos []models.Video
	if err := db.Find(&videos).Error; err != nil {
		return c.Status(500).SendString("Erreur lors de la récupération des utilisateurs")
	}

	return c.JSON(videos)
}

// Récupérer un utilisateur par ID
func GetVideoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var video models.Video
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	if err := db.First(&video, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Utilisateur non trouvé")
	}

	return c.JSON(video)
}

// Mettre à jour un utilisateur
func UpdateVideo(c *fiber.Ctx) error {
	id := c.Params("id")
	var video models.Video
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	if err := db.First(&video, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Utilisateur non trouvé")
	}

	if err := c.BodyParser(&video); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erreur dans les données envoyées")
	}

	if err := db.Save(&video).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la mise à jour de l'utilisateur")
	}

	return c.JSON(video)
}

// Supprimer un utilisateur
func DeleteVideo(c *fiber.Ctx) error {
	id := c.Params("id")
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	if err := db.Delete(&models.Video{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la suppression de l'utilisateur")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
