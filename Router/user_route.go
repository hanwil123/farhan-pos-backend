package Router

import (
	"Farhan-Backend-POS/controllers/controllerRestApi"

	fiber2 "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber2.App) {
	// app.Get("/api/register", Controllers2.Register)
	app.Post("/api/register", controllerRestApi.Register)
	// app.Post("/api/login", Controllers2.Login)
	// app.Get("/api/login", Controllers2.Login)
	// app.Get("/api/users", Controllers2.User)
	// app.Post("/api/users", Controllers2.User)
	// app.Post("/api/logout", Controllers2.Logout)

}
