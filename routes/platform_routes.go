package routes

import (
	"github.com/abdealt/videomaker/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPlatformRoutes(app *fiber.App) {
	// Routes model Platform
	// Route pour créer une platforme
	app.Post("/platform", controllers.CreatePlatform)

	// Route pour récupérer tous les platforme
	app.Get("/platform", controllers.GetPlatforms)

	// Route pour récupérer une platforme par ID
	app.Get("/platform/:id", controllers.GetPlatformByID)

	// Route pour mettre à jour une platforme
	app.Put("/platform/:id", controllers.UpdatePlatform)

	// Route pour supprimer une platforme
	app.Delete("/platform/:id", controllers.DeletePlatform)
}
