package main

import (
	"api-gateway/controller"
	"api-gateway/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi from api gateway")
	})

	group := app.Group("/api")

	userController := controller.NewUserContoller("http://localhost:3001/user")
	employeeController := controller.NewEmployeeContoller("http://localhost:3002/employee")

	router.NewEmployeeRouter(group, employeeController)
	router.NewUserRouter(group, userController)

	port := 3000
	fmt.Printf("api gateway is running on :%d...\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting api gateway: %v\n", err)
	}
}
