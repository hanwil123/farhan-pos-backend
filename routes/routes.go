package routes

import (
	"Farhan-Backend-POS/controllers/controllerRestApi"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Category routes
	categoryGroup := app.Group("/api/categories")
	categoryGroup.Post("/", controllerRestApi.CreateCategoryControllersApi)
	categoryGroup.Get("/", controllerRestApi.GetCategorieControllerApi) // New route for listing all categories
	categoryGroup.Get("/:id", controllerRestApi.GetCategoryByIdControllerApi)

	// Product routes
	productGroup := app.Group("/api/products")
	productGroup.Post("/", controllerRestApi.CreateProductControllerApi)
}
