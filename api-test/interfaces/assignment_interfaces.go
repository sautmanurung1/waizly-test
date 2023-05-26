package interfaces

import (
	"api-test/models"

	"github.com/labstack/echo/v4"
)

type AssignmentRepository interface {
	CreateAssignmentRepository(assignment models.Assignment) error
	GetAssignmentByIdRepository(id int) (models.Assignment, error)
	GetAllAssignmentRepository() ([]models.Assignment, error)
	UpdateAssignmentRepository(id int, assignment models.Assignment) (models.Assignment, error)
	DeleteAssignmentRepository(id int, assign models.Assignment) error
}

type AssignmentService interface {
	CreateAssignmentService(assignment models.Assignment) (string, error)
	GetAssginmentByIdService(id int) (models.Assignment, error)
	GetAllAssignmentService() (assignment []models.Assignment, err error)
	UpdateAssignmentService(id int, assignment models.Assignment) (models.Assignment, error)
	DeleteAssignmentService(id int, assignment models.Assignment) (string, error)
}

type AssignmentHandler interface {
	CreateAssignment(c echo.Context) error
	GetAssignmentById(c echo.Context) error
	GetAllAssignment(c echo.Context) error
	UpdateAssignment(c echo.Context) error
	DeleteAssignment(c echo.Context) error
}
