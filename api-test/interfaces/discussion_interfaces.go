package interfaces

import (
	"api-test/models"

	"github.com/labstack/echo/v4"
)

type DiscussionsRepository interface {
	CreateDiscussionsRepository(discussions models.Discussions) error
	GetAllDiscussionsRepository() (discussions []models.Discussions, err error)
}

type DiscussionsService interface {
	CreateDiscussionsService(discussions models.Discussions) error
	GetAllDiscussionsService() (discussions []models.Discussions, err error)
}

type DiscussionsHandler interface {
	CreateDiscussions(c echo.Context) error
	GetAllDiscussions(c echo.Context) error
}
