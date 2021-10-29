package routes

import (
	"test/waldhalf/gofiber-tpl/src/controllers"
	"test/waldhalf/gofiber-tpl/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App){
	api := app.Group("api")
	admin := api.Group("admin")
	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)
	
	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("/user", controllers.User)
	adminAuthenticated.Get("/logout", controllers.Logout)
	adminAuthenticated.Put("/users/info", controllers.UpdateProfile)
	adminAuthenticated.Put("/users/password", controllers.UpdatePassword)
	adminAuthenticated.Get("/get-admins", controllers.GetAdmin)
	adminAuthenticated.Get("/ambassadors", controllers.Ambassadors)

	// POST
	adminAuthenticated.Get("/posts", controllers.Posts)
	adminAuthenticated.Post("/posts", controllers.CreatePosts)
	adminAuthenticated.Get("/posts/:id", controllers.GetPost)
	adminAuthenticated.Put("/posts/:id", controllers.UpdatePost)
	adminAuthenticated.Delete("/posts/:id", controllers.DeletePost)

	// PRODUCT
	adminAuthenticated.Get("/products", controllers.Products)
	adminAuthenticated.Post("/products", controllers.CreateProducts)
	adminAuthenticated.Get("/products/:id", controllers.GetProduct)
	adminAuthenticated.Put("/products/:id", controllers.UpdateProduct)
	adminAuthenticated.Delete("/products/:id", controllers.DeleteProduct)
	adminAuthenticated.Get("/users/:id/links", controllers.Link)
	
	// ORDER
	adminAuthenticated.Get("/orders", controllers.Orders)

	// AMBASSADOR
	ambassador := api.Group("ambassador")
	ambassador.Post("/register", controllers.Register)
	ambassador.Post("/login", controllers.Login)

	// Frontend/Backend
	ambassador.Get("/products/frontend", controllers.ProductsFrontend)
	ambassador.Get("/products/backend", controllers.ProductsBackend)

	ambassadorAuthenticated := ambassador.Use(middlewares.IsAuthenticated)
	ambassadorAuthenticated.Get("/user", controllers.User)
	ambassadorAuthenticated.Get("/logout", controllers.Logout)
	ambassadorAuthenticated.Put("/users/info", controllers.UpdateProfile)
	ambassadorAuthenticated.Put("/users/password", controllers.UpdatePassword)
	ambassadorAuthenticated.Post("/links", controllers.CreateLink)
	ambassadorAuthenticated.Get("/stats", controllers.Stats)
	ambassadorAuthenticated.Get("/rankings", controllers.Rankings)


}