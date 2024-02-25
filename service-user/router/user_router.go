package router

import (
	"fmt"
	"service-user/controller"
	"service-user/model"

	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	controller.UserController
}

func (user UserRouter) Login(c *fiber.Ctx) error {
	var requestBody model.User

	c.BodyParser(&requestBody)
	code, response := user.UserLogin(c, requestBody)
	fmt.Println(response)

	if code != 200 {
		c.Status(code)
		return c.JSON(response)
	}

	return c.JSON(response)
}

func (user UserRouter) CreateUser(c *fiber.Ctx) error {
	var requestBody model.User

	c.BodyParser(&requestBody)

	response := user.Register(c, requestBody)
	if response.Code != 200 {
		c.Status(response.Code)
		return c.JSON(response.Data)
	}

	return c.JSON(response)
}

func (user UserRouter) Auth(c *fiber.Ctx) error {
	err := user.GetUser(c)
	if err != nil {
		return err
	}

	return c.JSON("OK")
}

func NewUserRouter(r fiber.Router, us controller.UserController) {

	router := &UserRouter{us}
	r.Post("/register", router.CreateUser)
	r.Post("/login", router.Login)
	r.Get("/auth", router.Auth)

}
