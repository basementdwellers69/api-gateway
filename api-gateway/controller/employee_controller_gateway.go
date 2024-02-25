package controller

import (
	"api-gateway/model"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type EmployeeController struct {
	string
}

func (employee EmployeeController) NewEmployee(c *fiber.Ctx, payload []byte) (int, interface{}) {
	// Create a new HTTP request
	agent := fiber.Post(employee.string + "/new")
	agent.Body(payload)

	agent.Set("access_token", c.Get("access_token"))
	agent.Set("Content-Type", "application/json")

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return statusCode, fiber.Map{
			"error": body,
			"errs":  errs,
		}
	}
	defer agent.ConnectionClose()

	var res model.EmployeeResponse
	err := json.Unmarshal(body, &res)
	if err != nil {
		return fiber.StatusInternalServerError, err.Error()
	}

	return statusCode, res
}

func NewEmployeeContoller(base_uri string) EmployeeController {
	return EmployeeController{base_uri}
}
