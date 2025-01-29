package routes

import (
	"github.com/abdealt/videomaker/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupLoginRoutes(app *fiber.App) {

	app.Post("/login", controllers.LoginUserController)

}
