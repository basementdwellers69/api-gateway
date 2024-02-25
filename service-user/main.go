package main

import (
	"fmt"
	"service-user/config"
	"service-user/controller"
	"service-user/router"

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
		return c.SendString("Hi from service-user")
	})

	group := app.Group("/user")

	userController := controller.NewUserContoller(db)

	router.NewUserRouter(group, userController)

	port := 3001
	fmt.Printf("Service user is running on :%d...\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting Service user: %v\n", err)
	}
}
