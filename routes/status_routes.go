package routes

import (
	"github.com/abdealt/videomaker/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupStatusRoutes(app *fiber.App) {
	// Routes model status
	// Route pour créer un Status
	app.Post("/status", controllers.CreateStatus)

	// Route pour récupérer tous les Status
	app.Get("/status", controllers.GetStatus)

	// Route pour récupérer un Status par ID
	app.Get("/status/:id", controllers.GetStatusByID)

	// Route pour mettre à jour un Status
	app.Put("/status/:id", controllers.UpdateStatus)

	// Route pour supprimer un Status
	app.Delete("/status/:id", controllers.DeleteStatus)
}
