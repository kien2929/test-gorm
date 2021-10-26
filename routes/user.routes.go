package routes

import (
	Controllers "vap/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	router.Post("/register", Controllers.Register)
}
