package controllerRestApi

import (
	"Farhan-Backend-POS/client"
	"Farhan-Backend-POS/proto"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateCategoryControllersApi(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	responseCreateCategory, errCreateCategory := client.CategoryClient.CreateCategory(ctx, &proto.CategoryRequest{
		Name: data["name"],
	})
	if errCreateCategory != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid create category: " + errCreateCategory.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"Name": responseCreateCategory.Name,
	})
}
