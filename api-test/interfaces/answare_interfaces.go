package interfaces

import (
	"api-test/models"

	"github.com/labstack/echo/v4"
)

type AnswareRepository interface {
	CreateAnsware(answare models.Answare) error
	GetAnswareById(id int) (models.Answare, error)
	GetAllAnsware() (ans []models.Answare, err error)
	UpdateAnsware(id int, answare models.Answare) (models.Answare, error)
	DeleteAnsware(id int, ans models.Answare) error
}

type AnswareService interface {
	CreateAnswareService(answare models.Answare) (string, error)
	GetAnswareByIdService(id int) (models.Answare, error)
	GetAllAnswareService() (ans []models.Answare, err error)
	UpdateAnswareService(id int, answare models.Answare) (models.Answare, error)
	DeleteAnswareService(id int, answare models.Answare) (string, error)
}

type AnswareHandler interface {
	CreateAnsware(c echo.Context) error
	GetAnswareById(c echo.Context) error
	GetAllAnsware(c echo.Context) error
	UpdateAnsware(c echo.Context) error
	DeleteAnsware(c echo.Context) error
}
