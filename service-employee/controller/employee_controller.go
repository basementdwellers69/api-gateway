package controller

import (
	"service-employee/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmployeeController struct {
	*gorm.DB
}

func (db EmployeeController) NewEmployee(c *fiber.Ctx, request model.Employee) model.WebResponse {

	request.ID = uuid.New().String()

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
		Data:   request,
	}
}

func NewEmployeeContoller(client *gorm.DB) EmployeeController {
	return EmployeeController{client}
}
