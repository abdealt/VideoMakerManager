package routes

import (
	"videomaker/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupStatusRoutes(router fiber.Router) {
	statuses := router.Group("/statuses")
	// Routes model status
	// Route pour créer un Status
	statuses.Post("/", controllers.CreateStatus)

	// Route pour récupérer tous les Status
	statuses.Get("/", controllers.GetStatus)

	// Route pour récupérer un Status par ID
	statuses.Get("/:id", controllers.GetStatusByID)

	// Route pour mettre à jour un Status
	statuses.Put("/:id", controllers.UpdateStatus)

	// Route pour supprimer un Status
	statuses.Delete("/:id", controllers.DeleteStatus)
}
