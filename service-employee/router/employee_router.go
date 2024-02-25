package router

import (
	"service-employee/controller"
	"service-employee/middleware"
	"service-employee/model"

	"github.com/gofiber/fiber/v2"
)

type EmployeeRouter struct {
	controller.EmployeeController
}

func (employee EmployeeRouter) GetEmployee(c *fiber.Ctx) error {
	return c.SendString("Hi from service-employee")
}

func (employee EmployeeRouter) CreateEmployee(c *fiber.Ctx) error {
	var requestBody model.Employee

	c.BodyParser(&requestBody)

	response := employee.NewEmployee(c, requestBody)

	if response.Code != 200 {
		c.Status(response.Code)
		return c.JSON(response.Data)
	}

	return c.JSON(response)
}

func NewEmployeeRouter(r fiber.Router, us controller.EmployeeController) {

	router := &EmployeeRouter{us}
	r.Get("/user", router.GetEmployee)
	r.Post("/new", middleware.Authentication, router.CreateEmployee)

}
