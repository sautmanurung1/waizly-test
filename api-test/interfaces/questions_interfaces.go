package interfaces

import (
	"api-test/models"

	"github.com/labstack/echo/v4"
)

type QuestionRepository interface {
	CreateQuestionRepository(question models.Question) error
	GetQuestionByIdRepository(id int) (models.Question, error)
	GetAllQuestionRepository() ([]models.Question, error)
	UpdateQuestionRepository(id int, question models.Question) (models.Question, error)
	DeleteQuestionRepository(id int, question models.Question) error
}

type QuestionService interface {
	CreateQuestionService(question models.Question) (string, error)
	GetQuestionByIdService(id int) (models.Question, error)
	GetAllQuestionService() (question []models.Question, err error)
	UpdateQuestionService(id int, question models.Question) (models.Question, error)
	DeleteQuestionService(id int, question models.Question) (string, error)
}

type QuestionHandler interface {
	CreateQuestion(c echo.Context) error
	GetQuestionById(c echo.Context) error
	GetAllQuestion(c echo.Context) error
	UpdateQuestion(c echo.Context) error
	DeleteQuestion(c echo.Context) error
}
