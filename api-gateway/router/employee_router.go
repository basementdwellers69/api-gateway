package router

import (
	"api-gateway/controller"
	"api-gateway/middleware"
	"api-gateway/model"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type EmployeeRouter struct {
	controller.EmployeeController
}

func (employee *EmployeeRouter) CreateEmployee(c *fiber.Ctx) error {
	var bodyRequest model.EmployeeBodyReq

	c.BodyParser(&bodyRequest)

	payload, err := json.Marshal(bodyRequest)
	if err != nil {
		c.Status(500)
		c.SendString(err.Error())
	}

	code, resp := employee.NewEmployee(c, payload)
	if code != 200 {
		c.Status(code)
		c.JSON(resp)
	}
	return c.JSON(resp)
}

func NewEmployeeRouter(r fiber.Router, us controller.EmployeeController) {

	router := &EmployeeRouter{us}
	r.Post("/employee", middleware.Authorization, router.CreateEmployee)

}
