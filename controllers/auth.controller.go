package controllers

import (
	"vap/db"
	"github.com/gofiber/fiber/v2"
)
func Register(c *fiber.Ctx) error {
	user := c.Body()
	db.DB.Create(&user)
	return c.JSON(fiber.Map{
		"success": true,
	})
}
