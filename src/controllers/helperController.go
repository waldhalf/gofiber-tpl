package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func IsResultHasRows(c *fiber.Ctx, result *gorm.DB)(*fiber.Ctx){
	if result.RowsAffected <= 0 {
		return c.Status(204)
	}
	return nil
}