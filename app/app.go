package app

import (
	"videomaker/database"
	"videomaker/pkg/migrations"
	"videomaker/routes"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	// Migration de la base de données
	if err := migrations.MigrateAll(db); err != nil {
		panic(err)
	}

	// Instance de l'application avec fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler, // On ajoutera un gestionnaire d'erreur personnalisé
	})

	// Groupe de routes pour l'API v1
	v1 := app.Group("/api/v1")

	// Initialisation des routes
	routes.SetupAuthRoutes(v1)
	routes.SetupPlatformRoutes(v1)
	routes.SetupUserRoutes(v1)
	routes.SetupStatusRoutes(v1)
	routes.SetupVideoRoutes(v1)

	// Lancement de l'application sur le port 3000
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}

// Gestionnaire d'erreur personnalisé
func customErrorHandler(c *fiber.Ctx, err error) error {
	// Code par défaut est 500
	code := fiber.StatusInternalServerError

	// Vérifie si c'est une erreur Fiber
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Renvoie une réponse JSON
	return c.Status(code).JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
	})
}
