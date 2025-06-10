package main

import (
	"Farhan-Backend-POS/Router"
	"Farhan-Backend-POS/client"
	"Farhan-Backend-POS/database"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	client.InitializeClient()
	database.ConnectUser()
	database.ConnectCategory()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     "*",
	}))

	// Router.SetupBookmarks(app)
	Router.Setup(app)
	var port = envPortOr("3000")
	err := app.Listen(port)
	if err != nil {
		panic(err)
	}

}
func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
