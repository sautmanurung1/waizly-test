package controller

import (
	"api-test/interfaces"
	"api-test/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type assignmentController struct {
	s interfaces.AssignmentService
}

func NewAssignmentController(assignmentService interfaces.AssignmentService) interfaces.AssignmentHandler {
	return &assignmentController{
		s: assignmentService,
	}
}

func (h *assignmentController) CreateAssignment(c echo.Context) error {
	assign := models.Assignment{}

	e := c.Bind(&assign)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "error to create assignment",
			"Data":  e.Error(),
		})
	}

	result, err := h.s.CreateAssignmentService(assign)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": "error to create assignment",
			"Data":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to create assignment",
		"Data":    result,
	})
}

func (h *assignmentController) GetAssignmentById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	assign, e := h.s.GetAssginmentByIdService(id)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "error to get the Assignment by id",
			"Error":   e.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success get assignment by id",
		"Data":    assign,
	})
}

func (h *assignmentController) GetAllAssignment(c echo.Context) error {
	assignments, err := h.s.GetAllAssignmentService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "error to get all assignment",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success get all assignment",
		"Data":    assignments,
	})
}

func (h *assignmentController) UpdateAssignment(c echo.Context) error {
	var assignments models.Assignment
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&assignments); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "error to update the assignment",
			"Error":   err.Error(),
		})
	}

	result, e := h.s.UpdateAssignmentService(id, assignments)

	if e != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "error to update the assignment",
			"Error":   e.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to update assignment",
		"Data":    result,
	})
}

func (h *assignmentController) DeleteAssignment(c echo.Context) error {
	var assign models.Assignment
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&assign); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "error to delete the assignment",
			"Data":  err.Error(),
		})
	}

	assignments, er := h.s.DeleteAssignmentService(id, assign)

	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": "error To delete the assignment",
			"Data":  er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to delete the assignment",
		"Data":    assignments,
	})
}
