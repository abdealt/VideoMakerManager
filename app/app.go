package app

import (
	"github.com/abdealt/videomaker/database"
	"github.com/abdealt/videomaker/pkg/migrations"
	"github.com/abdealt/videomaker/routes"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	// Connexion a la base de données
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	// Migration de la base de données
	if err := migrations.MigrateAll(db); err != nil {
		panic(err)
	}

	//intance de l'application avec fiber
	app := fiber.New()

	// Initialisaiton des routes
	routes.SetupLoginRoutes(app)
	routes.SetupPlatformRoutes(app)
	routes.SetupUserRoutes(app)
	routes.SetupStatusRoutes(app)
	routes.SetupVideoRoutes(app)

	// Lancement de l'application sur le port 3000
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
