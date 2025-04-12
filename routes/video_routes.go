package routes

import (
	"videomaker/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupVideoRoutes(router fiber.Router) {
	videos := router.Group("/videos")
	// Routes model Video
	// Route pour créer une video
	videos.Post("/video", controllers.CreateVideo)

	// Route pour récupérer tous les video
	videos.Get("/video", controllers.GetVideos)

	// Route pour récupérer une video par ID
	videos.Get("/video/:id", controllers.GetVideoByID)

	// Route pour mettre à jour une video
	videos.Put("/video/:id", controllers.UpdateVideo)

	// Route pour supprimer une video
	videos.Delete("/video/:id", controllers.DeleteVideo)
}
