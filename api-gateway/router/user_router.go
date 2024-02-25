package router

import (
	"api-gateway/controller"
	"api-gateway/model"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	controller.UserController
}

func (user *UserRouter) UserTest(c *fiber.Ctx) error {
	return user.Test(c)
}

func (user *UserRouter) Login(c *fiber.Ctx) error {
	var bodyRequest model.UserBodyReq
	c.BodyParser(&bodyRequest)

	payload, err := json.Marshal(bodyRequest)
	if err != nil {
		c.Status(500)
		c.SendString(err.Error())
	}

	code, resp := user.UserLogin(c, payload)
	if code != 200 {
		fmt.Println(code, resp)

		c.Status(code)
		c.JSON(resp)
	}

	return c.JSON(resp)
}

func NewUserRouter(r fiber.Router, us controller.UserController) {

	router := &UserRouter{us}
	r.Get("/user", router.UserTest)
	r.Post("/login", router.Login)

}
