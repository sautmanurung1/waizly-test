package controller

import (
	"api-test/interfaces"
	"api-test/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type questionController struct {
	s interfaces.QuestionService
}

func NewQuestionController(questionService interfaces.QuestionService) interfaces.QuestionHandler {
	return &questionController{
		s: questionService,
	}
}

func (h *questionController) CreateQuestion(c echo.Context) error {
	questions := models.Question{}

	e := c.Bind(&questions)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": e.Error(),
		})
	}

	question, err := h.s.CreateQuestionService(questions)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"Error":   "Error to Create Question",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "Success Create Question",
		"Data":    question,
	})
}

func (h *questionController) GetQuestionById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	questions, er := h.s.GetQuestionByIdService(id)
	if er != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  http.StatusBadRequest,
			"Message": "Error To Get The Questions",
			"Error":   er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success Get Question By ID",
		"Data":    questions,
	})
}

func (h *questionController) GetAllQuestion(c echo.Context) error {
	questions, err := h.s.GetAllQuestionService()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  http.StatusBadRequest,
			"Message": "Error TO Get All Questions",
			"Data":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success Get All Questions",
		"Data":    questions,
	})
}

func (h *questionController) UpdateQuestion(c echo.Context) error {
	var questions models.Question
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&questions); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  http.StatusBadRequest,
			"Message": "Error",
			"Error":   err.Error(),
		})
	}

	result, er := h.s.UpdateQuestionService(id, questions)

	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  http.StatusInternalServerError,
			"Message": "Error To Update The Questions",
			"Data":    er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  http.StatusOK,
		"Message": "Success Update Questions",
		"Data":    result,
	})
}

func (h *questionController) DeleteQuestion(c echo.Context) error {
	var questions models.Question
	id, _ := strconv.Atoi(c.Param("id"))

	result, er := h.s.DeleteQuestionService(id, questions)

	if er != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": "error to delete the assignment",
			"Data":  er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success to delete the assignment",
		"Data":    result,
	})
}
