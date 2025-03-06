package pages

import "github.com/gofiber/fiber/v2"

func Signup(c *fiber.Ctx) error {
	return c.SendFile("./public/html/signup.html")
}
