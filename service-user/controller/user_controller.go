package controller

import (
	"errors"
	"fmt"
	"net/mail"
	"service-user/helpers"
	"service-user/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserController struct {
	*gorm.DB
}

func (db UserController) Register(c *fiber.Ctx, request model.User) model.WebResponse {

	request.ID = uuid.New().String()
	request.Password = helpers.HashPassword([]byte(request.Password))

	_, emailErr := mail.ParseAddress(request.Email)

	if emailErr != nil {
		return model.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   errors.New("invalid email").Error(),
		}
	}

	err := db.Create(&request).Error

	if err != nil {
		return model.WebResponse{
			Code:   500,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		}
	}

	return model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   request.Email,
	}

}

func (db UserController) GetUser(c *fiber.Ctx) error {
	access_token := c.Get("access_token")
	var user model.User

	if len(access_token) == 0 {
		return c.Status(401).SendString("Invalid token: Access token missing")
	}

	checkToken, err := helpers.VerifyToken(access_token)

	if err != nil {
		return c.Status(401).SendString("Invalid token: Failed to verify token")
	}

	fmt.Println(checkToken, "CEKKKK", checkToken["email"])

	userErr := db.Where("email = ?", checkToken["email"]).First(&user).Error
	if userErr != nil {
		return c.Status(401).SendString("Invalid token: User not found")
	}

	c.Locals("user", user)

	return nil
}

func (db UserController) UserLogin(c *fiber.Ctx, request model.User) (int, interface{}) {
	var user model.User
	result := db.First(&user, "email = ?", request.Email)

	if result.Error != nil {
		return 500, model.WebResponse{
			Code:   500,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   result.Error.Error(),
		}
	}

	checkPassword := helpers.ComparePassword([]byte(user.Password), []byte(request.Password))
	if !checkPassword {
		return 401, model.WebResponse{
			Code:   401,
			Status: "BAD_REQUEST",
			Data:   errors.New("invalid password").Error(),
		}
	}

	access_token := helpers.SignToken(request.Email)

	return 200, model.LoginResponse{
		Code:        200,
		Status:      "OK",
		AccessToken: access_token,
		Data:        user,
	}

}

func NewUserContoller(client *gorm.DB) UserController {
	return UserController{client}
}
