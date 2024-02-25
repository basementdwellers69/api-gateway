package main

import (
	"fmt"
	"service-employee/config"
	"service-employee/controller"
	"service-employee/router"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.InitPostgresDatabase()
	db = config.GetPostgresDatabase()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi from service-employee")
	})

	group := app.Group("/employee")

	employeeController := controller.NewEmployeeContoller(db)
	router.NewEmployeeRouter(group, employeeController)

	port := 3002
	fmt.Printf("Service employee is running on :%d...\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting Service employee: %v\n", err)
	}
}
