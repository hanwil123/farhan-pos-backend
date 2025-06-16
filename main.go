package grpcClient

import (
	"Farhan-Backend-POS/Router"
	grpcClient "Farhan-Backend-POS/cmd/grpc-client"
	"Farhan-Backend-POS/database"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	grpcClient.InitializedClient()
	database.ConnectUser()
	database.ConnectCategory()
	database.ConnectProduct()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     "*",
	}))

	// Router.SetupBookmarks(app)
	Router.Setup(app)
	Router.SetupRoutesProduct(app)
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
