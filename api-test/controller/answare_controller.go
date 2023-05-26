package controller

import (
	"api-test/interfaces"
	"api-test/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type answareController struct {
	s interfaces.AnswareService
}

func NewAnswareController(answareService interfaces.AnswareService) interfaces.AnswareHandler {
	return &answareController{
		s: answareService,
	}
}

func (a *answareController) CreateAnsware(c echo.Context) error {
	ans := models.Answare{}

	e := c.Bind(&ans)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "error to create answare",
			"Data":  e.Error(),
		})
	}

	answered, err := a.s.CreateAnswareService(ans)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": "error to create answare",
			"Data":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to create answare",
		"Data":    answered,
	})
}

func (a *answareController) GetAnswareById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ans, e := a.s.GetAnswareByIdService(id)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "error to get answare by id",
			"Error":   e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to get answare by id",
		"Data":    ans,
	})
}

func (a *answareController) GetAllAnsware(c echo.Context) error {
	ans, err := a.s.GetAllAnswareService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "error to get all answare",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to get all answare",
		"Data":    ans,
	})
}

func (a *answareController) UpdateAnsware(c echo.Context) error {
	var ans models.Answare
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&ans); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "error to update the answare",
			"Error":   err.Error(),
		})
	}

	result, e := a.s.UpdateAnswareService(id, ans)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "error to update the answare",
			"Error":   e.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success update answare",
		"Data":    result,
	})
}

func (a *answareController) DeleteAnsware(c echo.Context) error {
	var ans models.Answare
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&ans); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "error to delete the answare",
			"Data":  err.Error(),
		})
	}

	result, er := a.s.DeleteAnswareService(id, ans)

	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": "error to delete the answare",
			"Data":  er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to delete the answare",
		"Data":    result,
	})
}
