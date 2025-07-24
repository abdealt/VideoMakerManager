package routes

import (
	"videomaker/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupVideoRoutes(router fiber.Router) {
	videos := router.Group("/videos")

	// Routes model Video
	// Route pour créer une vidéo
	videos.Post("/", controllers.CreateVideo)

	// Route pour récupérer toutes les vidéos
	videos.Get("/", controllers.GetVideos)

	// Route pour récupérer une vidéo par ID
	videos.Get("/:id", controllers.GetVideoByID)

	// Route pour mettre à jour une vidéo
	videos.Put("/:id", controllers.UpdateVideo)

	// Route pour supprimer une vidéo
	videos.Delete("/:id", controllers.DeleteVideo)
}
