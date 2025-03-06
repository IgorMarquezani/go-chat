package pages

import "github.com/gofiber/fiber/v2"

func Signin(c *fiber.Ctx) error {
	return c.SendFile("./public/html/signin.html")
}
