package controller

import (
	"api-test/interfaces"
	"api-test/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type discussionsController struct {
	s interfaces.DiscussionsService
}

func NewDiscussionsController(disucssionsService interfaces.DiscussionsService) interfaces.DiscussionsHandler {
	return &discussionsController{
		s: disucssionsService,
	}
}

func (h *discussionsController) CreateDiscussions(c echo.Context) error {
	discuss := models.Discussions{}

	e := c.Bind(&discuss)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "error to create assignment",
			"Data":  e.Error(),
		})
	}

	discussions := h.s.CreateDiscussionsService(discuss)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success create discussions",
		"Data":    discussions,
	})
}

func (h *discussionsController) GetAllDiscussions(c echo.Context) error {
	discuss, err := h.s.GetAllDiscussionsService()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "error to get all discussions",
			"Data":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success get all discussions",
		"Data":    discuss,
	})
}
