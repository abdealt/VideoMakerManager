package routes

import (
	"github.com/abdealt/videomaker/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	// Routes model user
	// Route pour créer un utilisateur
	app.Post("/users", controllers.CreateUser)

	// Route pour récupérer tous les utilisateurs
	app.Get("/users", controllers.GetUsers)

	// Route pour récupérer un utilisateur par ID
	app.Get("/users/:id", controllers.GetUserByID)

	// Route pour mettre à jour un utilisateur
	app.Put("/users/:id", controllers.UpdateUser)

	// Route pour supprimer un utilisateur
	app.Delete("/users/:id", controllers.DeleteUser)
}
