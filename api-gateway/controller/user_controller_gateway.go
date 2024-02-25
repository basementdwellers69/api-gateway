package controller

import (
	"api-gateway/model"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	string
}

func (user UserController) Test(c *fiber.Ctx) error {
	resp, err := http.Get(user.string + "/")

	if err != nil {
		panic(err)
	}

	resp.Header.Set("Accept", "application/json")

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return c.JSON(string(body))
}

func (user UserController) UserLogin(c *fiber.Ctx, payload []byte) (int, interface{}) {

	agent := fiber.Post(user.string + "/login")
	agent.Body(payload)

	agent.Set("Content-Type", "application/json")

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return statusCode, fiber.Map{
			"error": body,
			"errs":  errs,
		}
	}
	defer agent.ConnectionClose()

	// fmt.Println("Response Status:", resp.Status)
	// fmt.Println("Response Body:", string(body))

	var res model.LoginResponse

	err := json.Unmarshal(body, &res)
	if err != nil {
		return 500, err.Error()
	}

	return 200, res
}

func NewUserContoller(base_uri string) UserController {
	return UserController{base_uri}
}
