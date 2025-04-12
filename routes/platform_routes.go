package routes

import (
	"videomaker/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPlatformRoutes(router fiber.Router) {
	platform := router.Group("/platform")
	// Routes model Platform
	// Route pour créer une platforme
	platform.Post("/platform", controllers.CreatePlatform)

	// Route pour récupérer tous les platforme
	platform.Get("/platform", controllers.GetPlatforms)

	// Route pour récupérer une platforme par ID
	platform.Get("/platform/:id", controllers.GetPlatformByID)

	// Route pour mettre à jour une platforme
	platform.Put("/platform/:id", controllers.UpdatePlatform)

	// Route pour supprimer une platforme
	platform.Delete("/platform/:id", controllers.DeletePlatform)
}
