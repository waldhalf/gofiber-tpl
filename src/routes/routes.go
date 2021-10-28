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

	adminAuthenticated.Get("/posts", controllers.Posts)
	adminAuthenticated.Post("/posts", controllers.CreatePosts)
	adminAuthenticated.Get("/posts/:id", controllers.GetPost)
	adminAuthenticated.Put("/posts/:id", controllers.UpdatePost)
	adminAuthenticated.Delete("/posts/:id", controllers.DeletePost)
}