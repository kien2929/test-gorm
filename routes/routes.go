package routes

import (
	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(app *fiber.App) {
	user := app.Group("/user")
	UserRoutes(user)

	auth := app.Group("/auth")
	AuthRoutes(auth)

	link := app.Group("/link")
	LinkRoutes(link)
}
