package controllers

import (
	"fmt"
	"strconv"
	"test/waldhalf/gofiber-tpl/src/database"
	"test/waldhalf/gofiber-tpl/src/middlewares"
	"test/waldhalf/gofiber-tpl/src/models"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Validation pour le password
	if data["password"] != data["password_confirm"]{
		c.Status(400)
		return c.JSON(fiber.Map{
		"message": "passwords do not match",
		})
	}

	// Validation champs vides
	// TODO validation de l'email plus fine

	// VALIDATOR A METTRE DANS MODEL
	if 	len(data["first_name"]) == 0 || 
		len(data["last_name"]) == 0 ||
		len(data["email"]) == 0 ||
		len(data["password"]) < 8 ||
		len(data["password_confirm"]) < 8 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
		IsAdmin: false,
	}
	// On a attaché au user la méthode de "fabrication" du password
	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	// TODO Validation du form

	// On s'assure que le user est dans la table
	if user.Id== 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid credentials email",
		})
	}

	// On s'assure que le password match
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message":  "Invalid credentials password",
		})
	}

	// Generation du JWT
	payload := jwt.StandardClaims{
		Subject: strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message":  "Invalid credentials.",
			"error" : string(err.Error()),
		})
	}

	// Retour du token sous forme de cookie
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
			"message":  "Success login",
	})
}

func User(c *fiber.Ctx) error{
	id, _ := middlewares.GetUserId(c)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error{
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
			"message":  "Success logout",
	})
}

func UpdateProfile(c *fiber.Ctx) error{
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	id, _ := middlewares.GetUserId(c)

	user := models.User{
		Id:			id,
		FirstName: 	data["first_name"],
		LastName: 	data["last_name"],
		Email: 		data["email"],
	}
	fmt.Println(user)

	database.DB.Model(&user).Updates(&user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error{
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Validation pour le password
	if data["password"] != data["password_confirm"]{
		c.Status(400)
		return c.JSON(fiber.Map{
		"message": "passwords do not match",
		})
	}

	id, _ := middlewares.GetUserId(c)

	user := models.User{
		Id:			id,
	}
	
	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(&user)

	return c.JSON(user)
}