package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Authorization(c *fiber.Ctx) error {
	// Set the Content-Type header

	access_token := c.Get("access_token")
	if len(access_token) == 0 {
		return c.Status(401).SendString("Invalid token: Access token missing")
	}

	agent := fiber.Get("http://localhost:3001/user/auth")

	agent.Set("access_token", access_token)
	agent.Set("Content-Type", "application/json")

	statusCode, _, errs := agent.Bytes()
	if len(errs) > 0 || statusCode != 200 {
		return c.Status(401).SendString("Invalid token: Access token missing")
	}
	defer agent.ConnectionClose()

	// Add additional headers as needed
	//c.Set("access_token", access_token)

	return c.Next()
}
