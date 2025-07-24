package routes

import (
	"videomaker/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/login", controllers.LoginUserController)

}
