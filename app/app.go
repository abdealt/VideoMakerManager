// app/app.go
package app

import (
	"videomaker/database"
	"videomaker/pkg/auth"
	"videomaker/pkg/middleware"
	"videomaker/pkg/migrations"
	"videomaker/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() {
	// Connexion à la base de données
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	// Initialisation de JWT
	if err := auth.Init(); err != nil {
		panic(err)
	}

	// Migration de la base de données
	if err := migrations.MigrateAll(db); err != nil {
		panic(err)
	}

	// Instance de l'application avec fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	// Middleware global
	app.Use(recover.New())        // Récupération des panics
	app.Use(logger.New())         // Logs des requêtes
	app.Use(cors.New(cors.Config{ // Configuration CORS
		AllowOrigins:     "http://localhost:4200", // URL du frontend Angular
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// Groupe de routes pour l'API v1
	v1 := app.Group("/api/v1")

	// Routes publiques
	routes.SetupAuthRoutes(v1)

	// Routes protégées
	// Application du middleware d'authentification
	v1.Use(middleware.Protected())

	// Initialisation des routes protégées
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
