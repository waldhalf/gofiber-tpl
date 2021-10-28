package controllers

import (
	"test/waldhalf/gofiber-tpl/src/database"
	"test/waldhalf/gofiber-tpl/src/models"

	"github.com/gofiber/fiber/v2"
)

func GetAdmin(c *fiber.Ctx) error {
	var admins []models.User

	database.DB.Where("is_admin = true").Find(&admins)

	return c.JSON(admins)
}