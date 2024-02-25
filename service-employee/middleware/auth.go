package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var user_uri string = "http://localhost:3002/user"

func Authentication(c *fiber.Ctx) error {

	access_token := c.Get("access_token")

	if len(access_token) == 0 {
		return c.Status(401).SendString("Invalid token: Access token missing")
	}

	req, err := http.NewRequest("GET", user_uri+"/auth", nil)
	if err != nil {
		c.Status(500).SendString("Error creating request:" + err.Error())
	}

	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access_token", access_token)

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Status(500).SendString("Error sending request:" + err.Error())
	}
	defer resp.Body.Close()

	return c.Next()
}
