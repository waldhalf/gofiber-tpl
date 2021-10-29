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
		post := models.Post{
			Title 		: faker.Username(),
			SubTitle 	: faker.Username(),
			Content 	: faker.Paragraph(),
			AuthorId 	: uint(rand.Intn(90)+10),
			ImageUrl 	: faker.URL(),
			ReleaseDate	: faker.Date(),
			CategoryId  : uint(rand.Intn(90)+10),
		}
		database.DB.Create(&post)
	}
}

