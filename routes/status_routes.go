package routes

import (
	"videomaker/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupStatusRoutes(router fiber.Router) {
	status := router.Group("/status")
	// Routes model status
	// Route pour créer un Status
	status.Post("/status", controllers.CreateStatus)

	// Route pour récupérer tous les Status
	status.Get("/status", controllers.GetStatus)

	// Route pour récupérer un Status par ID
	status.Get("/status/:id", controllers.GetStatusByID)

	// Route pour mettre à jour un Status
	status.Put("/status/:id", controllers.UpdateStatus)

	// Route pour supprimer un Status
	status.Delete("/status/:id", controllers.DeleteStatus)
}
