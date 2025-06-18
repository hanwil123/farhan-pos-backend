// user api handler - FIXED VERSION
package restapiAuth

import (
	grpcClient "Farhan-Backend-POS/cmd/grpc-client"
	"Farhan-Backend-POS/proto"
	"context"
	"fmt"
	"time"

	fiber2 "github.com/gofiber/fiber/v2"
)

const SecretKey = "secret"

func Register(c *fiber2.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber2.StatusBadRequest).JSON(fiber2.Map{
			"message": "Invalid request body",
		})
	}
	fmt.Print("data register yang dimasukkan : ", data)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := grpcClient.UserClient.RegisterUser(ctx, &proto.RegisterUserRequest{
		Name:     data["name"],
		Email:    data["email"],
		Password: data["password"],
		Role:     data["role"],
	})
	if err != nil {
		return c.Status(fiber2.StatusInternalServerError).JSON(fiber2.Map{
			"message": "Invalid register: " + err.Error(),
		})
	}

	return c.JSON(fiber2.Map{
		"id":      resp.Id,
		"name":    resp.Name,
		"email":   resp.Email,
		"Role":    resp.Role,
		"message": resp.Message,
		"status":  resp.StatusCodeBerhasilRegister,
	})
}

func Login(c *fiber2.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber2.StatusBadRequest).JSON(fiber2.Map{
			"message": "Invalid request body",
		})
	}
	fmt.Print("data login yang dimasukkan : ", data)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	respLogin, errLogin := grpcClient.UserClient.LoginUser(ctx, &proto.LoginUserRequest{
		Email:    data["email"],
		Password: data["password"],
	})
	if errLogin != nil {
		return c.Status(fiber2.StatusInternalServerError).JSON(fiber2.Map{
			"message": "Login service error: " + errLogin.Error(),
		})
	}

	// Check status dari response
	if respLogin.StatusCodeBerhasilLogin != "200" {
		return c.Status(fiber2.StatusBadRequest).JSON(fiber2.Map{
			"message": respLogin.Message,
		})
	}

	// Set cookie dengan token (opsional)
	c.Cookie(&fiber2.Cookie{
		Name:     "jwt",
		Value:    respLogin.Token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})

	return c.JSON(fiber2.Map{
		"id":      respLogin.Id,
		"token":   respLogin.Token,
		"message": respLogin.Message,
		"status":  respLogin.StatusCodeBerhasilLogin,
	})
}
