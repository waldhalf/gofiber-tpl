package controllers

import (
	"test/waldhalf/gofiber-tpl/src/database"
	"test/waldhalf/gofiber-tpl/src/models"

	"github.com/gofiber/fiber/v2"
)

func Orders(c *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Preload("OrderItems").Find(&orders)
	for i, order := range orders {
		orders[i].Name = order.FullName()
		orders[i].Total = order.GetTotal()
	}
	return c.JSON(orders)
}

