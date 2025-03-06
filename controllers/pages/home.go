package pages

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.SendFile("./public/htlm/home.html")
}
