package restapiBakery

import (
	grpcClient "Farhan-Backend-POS/cmd/grpc-client"
	"Farhan-Backend-POS/modules/bakery/dto"
	"Farhan-Backend-POS/proto"
	"context"
	"strconv"
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

	responseCreateCategory, errCreateCategory := grpcClient.BakeryPOSClient.CreateCategory(ctx, &proto.CategoryRequest{
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

func CreateProductControllerApi(c *fiber.Ctx) error {
	var data dto.Product
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// price, err := strconv.ParseFloat(data["price"], 64)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "Invalid price format",
	// 	})
	// }

	// stockQuantity, err := strconv.ParseInt(data["stock_quantity"], 10, 32)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "Invalid stock quantity format",
	// 	})
	// }

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	responseAddProduct, errorAddProduct := grpcClient.BakeryPOSClient.CreateProduct(ctx, &proto.CreateProductRequest{
		Product: &proto.Product{
			Name:          data.Name,
			Description:   data.Description,
			Price:         data.Price,
			StockQuantity: int32(data.StockQuantity),
			CategoryId:    strconv.FormatUint(data.CategoryID, 10),
			ImageUrl:      data.ImageURL,
		},
	})
	if errorAddProduct != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid create product: " + errorAddProduct.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message":        "successful add product",
		"Id":             responseAddProduct.Product.Id,
		"Name":           responseAddProduct.Product.Name,
		"description":    responseAddProduct.Product.Description,
		"price":          responseAddProduct.Product.Price,
		"stock_quantity": responseAddProduct.Product.StockQuantity,
	})
}

func GetCategoryByIdControllerApi(c *fiber.Ctx) error {
	categoryId := c.Params("id")
	if categoryId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Category ID is required",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := grpcClient.BakeryPOSClient.GetCategoryById(ctx, &proto.GetCategoryByIdRequest{
		Id: categoryId,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get category: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id":   response.Id,
		"name": response.Name,
	})
}

func GetCategorieControllerApi(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	response, err := grpcClient.BakeryPOSClient.ListCategories(ctx, &proto.Empty{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get categories: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"categories": response.Categories,
	})
}

func UpdateProductControllerApi(c *fiber.Ctx) error {
	productId := c.Params("id")
	if productId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Product ID is required",
		})
	}
	var data dto.Product
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	responseUpdateProduct, err := grpcClient.BakeryPOSClient.UpdateProduct(ctx, &proto.UpdateProductRequest{
		Product: &proto.Product{
			Id:            productId,
			Name:          data.Name,
			Description:   data.Description,
			Price:         data.Price,
			StockQuantity: int32(data.StockQuantity),
			CategoryId:    strconv.FormatUint(data.CategoryID, 10),
			ImageUrl:      data.ImageURL,
		},
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update productttttttttt: " + err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"Id":             responseUpdateProduct.Product.Id,
		"Name":           responseUpdateProduct.Product.Name,
		"Description":    responseUpdateProduct.Product.Description,
		"Price":          responseUpdateProduct.Product.Price,
		"Stock Quantity": responseUpdateProduct.Product.StockQuantity,
		"Image Url":      responseUpdateProduct.Product.ImageUrl,
	})
}

func GetAllProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	responseProduct, err := grpcClient.BakeryPOSClient.ListProducts(ctx, &proto.Empty{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get product: " + err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"Product": responseProduct.Products,
	})
}

func DeleteProductControllerApi(c *fiber.Ctx) error {
	productId := c.Params("id")
	if productId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Product ID is required",
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	responseDeleteProduct, err := grpcClient.BakeryPOSClient.DeleteProduct(ctx, &proto.DeleteProductRequest{
		Id: productId,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete product: " + err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": responseDeleteProduct.MessageSuccesfull,
	})
}
