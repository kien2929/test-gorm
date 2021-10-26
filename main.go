package main

import (
	"vap/db"
	"vap/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.InitDB()
	app := fiber.New()
	routes.LoadRoutes(app)
	app.Listen(":3000")
}
