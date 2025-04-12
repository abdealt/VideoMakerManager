package controllers

import (
	"videomaker/database"
	"videomaker/models"

	"github.com/gofiber/fiber/v2"
)

// Créer un nouvel utilisateur
func CreatePlatform(c *fiber.Ctx) error {
	var platform models.Platform
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	// Enregistrer l'utilisateur dans la base de données avec Gorm
	if err := db.Create(&platform).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la création de l'utilisateur")
	}

	return c.Status(fiber.StatusCreated).JSON(platform)
}

func GetPlatforms(c *fiber.Ctx) error {
	// Récupère l'instance de la base de données depuis le contexte ou ailleurs
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).SendString("Erreur de connexion à la base de données")
	}

	var platforms []models.Platform
	if err := db.Find(&platforms).Error; err != nil {
		return c.Status(500).SendString("Erreur lors de la récupération des utilisateurs")
	}

	return c.JSON(platforms)
}

// Récupérer un utilisateur par ID
func GetPlatformByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var platform models.Platform
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	if err := db.First(&platform, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Utilisateur non trouvé")
	}

	return c.JSON(platform)
}

// Mettre à jour un utilisateur
func UpdatePlatform(c *fiber.Ctx) error {
	id := c.Params("id")
	var platform models.Platform
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	if err := db.First(&platform, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Utilisateur non trouvé")
	}

	if err := c.BodyParser(&platform); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erreur dans les données envoyées")
	}

	if err := db.Save(&platform).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la mise à jour de l'utilisateur")
	}

	return c.JSON(platform)
}

// Supprimer un utilisateur
func DeletePlatform(c *fiber.Ctx) error {
	id := c.Params("id")
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	if err := db.Delete(&models.Platform{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la suppression de l'utilisateur")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
