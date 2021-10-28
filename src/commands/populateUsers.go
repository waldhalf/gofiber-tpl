package main

import (
	"test/waldhalf/gofiber-tpl/src/database"
	"test/waldhalf/gofiber-tpl/src/models"

	"github.com/bxcodec/faker/v3"
)

func main(){
	database.Connect()
	for i := 0; i < 30; i++ {
		user := models.User{
			FirstName: faker.FirstName(),
			LastName: faker.LastName(),
			Email: faker.Email(),
			IsAdmin: false,
		}
		user.SetPassword("1234")
		database.DB.Create(&user)
	}
}

