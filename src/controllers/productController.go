package controllers

import (
	"context"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"test/waldhalf/gofiber-tpl/src/database"
	"test/waldhalf/gofiber-tpl/src/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Products(c *fiber.Ctx) error {
	var products []models.Product
	database.DB.Find(&products)
	return c.JSON(products)
}

func CreateProducts(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	database.DB.Create(&product)
	// On vide le cache Redis quand on update pour permettre un nouveau cache
	go deleteCache("products_backend")
	go deleteCache("products_frontend")

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	// TODO faire vérif si produit existe ou pas
	var product models.Product
	id, _ := strconv.Atoi(c.Params("id"))
	product.Id = uint(id)
	
	result := database.DB.Find(&product)
	IsResultHasRows(c, result)
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{}
	product.Id = uint(id)
	result := database.DB.Find(&product)

	IsResultHasRows(c, result)

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(&product)

	// On vide le cache Redis quand on update pour permettre un nouveau cache
	go deleteCache("products_backend")
	go deleteCache("products_frontend")

	return c.JSON(product)
}

func deleteCache(key string){
	database.Cache.Del(context.Background(), key)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{}
	product.Id = uint(id)
	result := database.DB.Find(&product)

	IsResultHasRows(c, result)

	database.DB.Delete(&product)
	
	// On vide le cache Redis quand on update pour permettre un nouveau cache
	go deleteCache("products_backend")
	go deleteCache("products_frontend")

	return c.JSON(fiber.Map{
			"message": "The product has been deleted!",
	})
}

func ProductsFrontend(c *fiber.Ctx)(error) {
	var products []models.Product
	// On cache les data dans Redis
	var ctx = context.Background()
	result, err := database.Cache.Get(ctx, "products_frontend").Result()
	if err != nil {
		database.DB.Find(&products)
		bytes, err := json.Marshal(products)
		if err != nil {
			panic(err)
		}
		if errKey := database.Cache.Set(ctx,"products_frontend", bytes, 30*time.Minute).Err(); errKey != nil {
			panic(errKey)
		}
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	return c.JSON(products)
}

func ProductsBackend(c *fiber.Ctx)(error) {
	var products []models.Product
	// On cache les data dans Redis
	var ctx = context.Background()
	result, err := database.Cache.Get(ctx, "products_backend").Result()
	if err != nil {
		database.DB.Find(&products)
		bytes, err := json.Marshal(products)
		if err != nil {
			panic(err)
		}
		database.Cache.Set(ctx,"products_backend", bytes, 30*time.Minute).Err()
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	var searchedProducts []models.Product
	if s := c.Query("s"); s != "" {
		lower := strings.ToLower(s)
		for _, product := range products {
			if 	strings.Contains(strings.ToLower(product.Title), lower) || 
				strings.Contains(strings.ToLower(product.Description), lower){
				searchedProducts = append(searchedProducts, product)
			}
		}
	} else{
		searchedProducts = products
	}
	// Sort by price
	if sortParam := c.Query("sort"); sortParam != "" {
		sortLower := strings.ToLower(sortParam)
		if sortLower == "asc" {
			sort.Slice(searchedProducts, func(i, j int) bool{
				return searchedProducts[i].Price < searchedProducts[j].Price 
			})
		}
		if sortLower == "desc" {
			sort.Slice(searchedProducts, func(i, j int) bool{
				return searchedProducts[i].Price > searchedProducts[j].Price 
			})
		}
	}

	// Pagination
	var total = len(searchedProducts)
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage := 9
	var data []models.Product

	if total <= page * perPage && total >= (page - 1) * perPage{
		data = searchedProducts[(page - 1) * perPage: total]
	} else if total >= page * perPage{
		data = searchedProducts[(page-1)*perPage : page*perPage]
	} else {
		data = []models.Product{}
	}


	return c.JSON(fiber.Map{
		"data" : 		data,
		"total" : 		total,
		"page": 		page,
		"last_page": 	total/perPage + 1,
	})
	// return c.JSON(data[(page-1)*perPage : page*perPage])
}