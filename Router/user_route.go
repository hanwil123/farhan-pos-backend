package Router

import (
	restapiAuth "Farhan-Backend-POS/modules/auth/delivery-handler/rest-api"
	restapiBakery "Farhan-Backend-POS/modules/bakery/delivery-handler/rest-api"

	fiber2 "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber2.App) {
	// app.Get("/api/register", Controllers2.Register)
	app.Post("/api/register", restapiAuth.Register)
	app.Post("/api/loginuser", restapiAuth.Login)
	app.Post("/api/create/category", restapiBakery.CreateCategoryControllersApi)
	// app.Get("/api/login", Controllers2.Login)
	// app.Get("/api/users", Controllers2.User)
	// app.Post("/api/users", Controllers2.User)
	// app.Post("/api/logout", Controllers2.Logout)

}
func SetupRoutesProduct(app *fiber2.App) {
	// Category routes
	categoryGroup := app.Group("/api/categories")
	categoryGroup.Post("/", restapiBakery.CreateCategoryControllersApi)
	// categoryGroup.Get("/:id", restapiBakery.GetCategoryByIdControllerApi)
	categoryGroup.Get("/allCategories", restapiBakery.GetCategorieControllerApi)

	// Product routes
	productGroup := app.Group("/api/products")
	productGroup.Post("/", restapiBakery.CreateProductControllerApi)
}
