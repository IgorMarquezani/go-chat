package routes

import (
	"app/core/controllers/pages"
	"app/core/controllers/users"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", pages.Signin)
	app.Get("/sign-up", pages.Signup)
	app.Get("/home", pages.Home)

	app.Post("/users/sign-up", users.Signup)

	app.Static("/static", "./public", fiber.Static{
		Browse: true,
	})
}
