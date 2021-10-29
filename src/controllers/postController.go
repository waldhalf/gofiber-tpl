package controllers

import (
	"strconv"
	"test/waldhalf/gofiber-tpl/src/database"
	"test/waldhalf/gofiber-tpl/src/models"

	"github.com/gofiber/fiber/v2"
)

func Posts(c *fiber.Ctx) error {
	var posts []models.Post
	database.DB.Find(&posts)
	return c.JSON(posts)
}

func CreatePosts(c *fiber.Ctx) error {
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return err
	}

	database.DB.Create(&post)
	return c.JSON(post)

}

func GetPost(c *fiber.Ctx) error {
	var post models.Post
	id, _ := strconv.Atoi(c.Params("id"))
	post.Id = uint(id)
	// database.DB.Where("id = ?", id).First(&post)
	result := database.DB.Find(&post)

	if result.RowsAffected <= 0 {
		c.Status(204)
	}

	return c.JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	post := models.Post{}
	post.Id = uint(id)
	if err := c.BodyParser(&post); err != nil {
		return err
	}

	oldPost := database.DB.Find(&post)

	if oldPost.RowsAffected <= 0 {
		c.Status(204)
	}

	newPost := database.DB.Model(&post).Updates(&post)

	return c.JSON(newPost)
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	post := models.Post{}
	post.Id = uint(id)
	oldPost := database.DB.Find(&post)

	if oldPost.RowsAffected <= 0 {
		c.Status(204)
		return nil
	}

	database.DB.Delete(&post)
	c.Status(200)
	return c.JSON(fiber.Map{
			"message": "The post has been deleted!",
	})
}