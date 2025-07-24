package routes

import (
	"videomaker/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupPlatformRoutes configure les routes pour le modèle Platform
// Cette fonction crée un groupe de routes pour les opérations CRUD sur les plateformes.
func SetupPlatformRoutes(router fiber.Router) {
	platforms := router.Group("/platforms")

	// Routes model Platform
	// Route pour créer une platforme
	platforms.Post("/", controllers.CreatePlatform)

	// Route pour récupérer tous les platforme
	platforms.Get("/", controllers.GetPlatforms)

	// Route pour récupérer une platforme par ID
	platforms.Get("/:id", controllers.GetPlatformByID)

	// Route pour mettre à jour une platforme
	platforms.Put("/:id", controllers.UpdatePlatform)

	// Route pour supprimer une platforme
	platforms.Delete("/:id", controllers.DeletePlatform)
}
