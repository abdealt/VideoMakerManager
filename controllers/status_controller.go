package controllers

import (
	"github.com/abdealt/videomaker/database"
	"github.com/abdealt/videomaker/models"

	"github.com/gofiber/fiber/v2"
)

// Créer un nouvel utilisateur
func CreateStatus(c *fiber.Ctx) error {
	var status models.Status
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	// Enregistrer l'utilisateur dans la base de données avec Gorm
	if err := db.Create(&status).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la création de l'utilisateur")
	}

	return c.Status(fiber.StatusCreated).JSON(status)
}

func GetStatus(c *fiber.Ctx) error {
	// Récupère l'instance de la base de données depuis le contexte ou ailleurs
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).SendString("Erreur de connexion à la base de données")
	}

	var statuss []models.Status
	if err := db.Find(&statuss).Error; err != nil {
		return c.Status(500).SendString("Erreur lors de la récupération des utilisateurs")
	}

	return c.JSON(statuss)
}

// Récupérer un utilisateur par ID
func GetStatusByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var status models.Status
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	if err := db.First(&status, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Utilisateur non trouvé")
	}

	return c.JSON(status)
}

// Mettre à jour un utilisateur
func UpdateStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var status models.Status
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	if err := db.First(&status, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Utilisateur non trouvé")
	}

	if err := c.BodyParser(&status); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erreur dans les données envoyées")
	}

	if err := db.Save(&status).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la mise à jour de l'utilisateur")
	}

	return c.JSON(status)
}

// Supprimer un utilisateur
func DeleteStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	if err := db.Delete(&models.Status{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la suppression de l'utilisateur")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
