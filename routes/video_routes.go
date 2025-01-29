package routes

import (
	"github.com/abdealt/videomaker/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupVideoRoutes(app *fiber.App) {
	// Routes model Video
	// Route pour créer une video
	app.Post("/video", controllers.CreateVideo)

	// Route pour récupérer tous les video
	app.Get("/video", controllers.GetVideos)

	// Route pour récupérer une video par ID
	app.Get("/video/:id", controllers.GetVideoByID)

	// Route pour mettre à jour une video
	app.Put("/video/:id", controllers.UpdateVideo)

	// Route pour supprimer une video
	app.Delete("/video/:id", controllers.DeleteVideo)
}
