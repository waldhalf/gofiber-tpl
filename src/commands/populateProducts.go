package main

import (
	"test/waldhalf/gofiber-tpl/src/database"
	"test/waldhalf/gofiber-tpl/src/models"

	"github.com/bxcodec/faker/v3"

	"math/rand"
)

func main(){
	database.Connect()
	for i := 0; i < 30; i++ {
		product := models.Product{
			Title 		:faker.Username(),
			Description : faker.Paragraph() ,
			Image: faker.URL(),
			Price 	:float64(rand.Intn(1000)),
		}
		database.DB.Create(&product)
	}
}

