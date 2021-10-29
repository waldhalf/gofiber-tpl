package controllers

import (
	"strconv"
	"test/waldhalf/gofiber-tpl/src/database"
	"test/waldhalf/gofiber-tpl/src/middlewares"
	"test/waldhalf/gofiber-tpl/src/models"

	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
)

type CreateLinkRequest struct {
	Products []int
}

func Link(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var links []models.Link
	database.DB.Where("user_id = ?", id).Find(&links)

	for i, link := range links {
		var orders []models.Order
		database.DB.Where("code = ? and complete = true", link.Code).Find(&orders)
		links[i].Orders = orders
	}
	return c.JSON(links)
}

func CreateLink(c *fiber.Ctx) error {
	var request CreateLinkRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	id, _ := middlewares.GetUserId(c)

	link := models.Link{
		UserId: id,
		Code: faker.Username(),
	}
	for _, productId := range request.Products{
		product := models.Product{}
		product.Id = uint(productId)
		link.Products = append(link.Products, product)
	}

	database.DB.Create(&link)
	return c.JSON(link)

}

func Stats(c *fiber.Ctx) error {
	id, _ := middlewares.GetUserId(c)
	var links []models.Link
	database.DB.Find(&links, models.Link{
		UserId: id,
	})
	var result []interface{}

	var orders []models.Order

	for _, link := range links {
		database.DB.Preload("OrdersItems").Find(&orders, &models.Order{
			Code: link.Code,
			Complete: true,
		})
		revenue := 0.0

		for _, order := range orders {
			revenue += order.GetTotal()
		}

		result = append(result, fiber.Map{
			"code" : link.Code,
			"count": len(orders),
			"revenu": revenue,
		})
	}
	return c.JSON(result)
}

func Rankings(c *fiber.Ctx) error {
	var users []models.User
	
	database.DB.Find(&users, models.User{
		IsAmbassador: true,
	})

	var result []interface{}
	
	for _, user := range users {
		ambassador :=  models.Ambassador(user)
		ambassador.CalculateRenevue(database.DB)
		result = append(result, fiber.Map{
			user.Name(): ambassador.Revenue,
		})
	}
	return c.JSON(result)
}
