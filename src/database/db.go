package database

import (
	"test/waldhalf/gofiber-tpl/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	var err error
	dsn := "host=db user=gorm password=gorm dbname=db_go port=5432 sslmode=disable TimeZone=Europe/Paris"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{}, models.Post{})
}