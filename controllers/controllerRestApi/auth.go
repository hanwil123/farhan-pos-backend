package controllerRestApi

import (
	"Farhan-Backend-POS/client"
	"Farhan-Backend-POS/proto"
	"context"
	"fmt"
	"math/rand"

	// "strconv"
	// "strings"
	"time"

	// "github.com/dgrijalva/jwt-go"
	fiber2 "github.com/gofiber/fiber/v2"
)

const SecretKey = "secret"
const SecretKey2 = "secret"

func Register(c *fiber2.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	rand.NewSource(time.Now().UnixNano())
	// randomUserID := uint(rand.Intn(10000) + 1)
	user, errRegister := client.UserClient.RegisterUser(ctx, &proto.RegisterUserRequest{
		Name:     data["name"],
		Email:    data["email"],
		Password: data["password"],
	})
	fmt.Printf("data user yang teregister : ", user)
	if errRegister != nil {
		c.Status(fiber2.StatusInternalServerError)
		return c.JSON(fiber2.Map{
			"message": "invalid register",
		})
	}
	return c.JSON(user)
}
