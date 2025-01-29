package controllers

import (
	"github.com/abdealt/videomaker/database"
	"github.com/abdealt/videomaker/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

// Créer un nouvel utilisateur
// Créer un nouvel utilisateur
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	// Enregistrer l'utilisateur dans la base de données avec Gorm
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erreur dans les données envoyées")
	}

	// Hacher le mot de passe avec bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors du hachage du mot de passe")
	}

	// Remplacer le mot de passe par le mot de passe haché
	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la création de l'utilisateur")
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	// Récupère l'instance de la base de données depuis le contexte ou ailleurs
	db, err := database.Connect()
	if err != nil {
		return c.Status(500).SendString("Erreur de connexion à la base de données")
	}

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return c.Status(500).SendString("Erreur lors de la récupération des utilisateurs")
	}

	return c.JSON(users)
}

// Récupérer un utilisateur par ID
func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	if err := db.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Utilisateur non trouvé")
	}

	return c.JSON(user)
}

// Mettre à jour un utilisateur
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	if err := db.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Utilisateur non trouvé")
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erreur dans les données envoyées")
	}

	if err := db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la mise à jour de l'utilisateur")
	}

	return c.JSON(user)
}

// Supprimer un utilisateur
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de l'accès à la base de données")
	}

	if err := db.Delete(&models.User{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erreur lors de la suppression de l'utilisateur")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
