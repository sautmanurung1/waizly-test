package routes

import (
	"api-test/infrastructure/database"

	"github.com/labstack/echo/v4"
)

func RootRoutes(echo *echo.Echo, conf database.Config) {
	RoutesAuth(echo, conf)
	RoutesAnsware(echo, conf)
	RoutesAssignment(echo, conf)
	RoutesDiscussions(echo, conf)
	RoutesQuestion(echo, conf)
}
