package routes

import (
	"videomaker/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	users := router.Group("/users")

	// Route pour créer un utilisateur
	users.Post("/", controllers.CreateUser)

	// Route pour récupérer tous les utilisateurs avec pagination
	users.Get("/", controllers.GetUsers)

	// Route pour récupérer un utilisateur par ID
	users.Get("/:id", controllers.GetUserByID)

	// Route pour mettre à jour un utilisateur
	users.Put("/:id", controllers.UpdateUser)

	// Route pour supprimer un utilisateur
	users.Delete("/:id", controllers.DeleteUser)
}
