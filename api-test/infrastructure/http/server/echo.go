package server

import (
	"api-test/infrastructure/database"
	"api-test/infrastructure/http/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server() *echo.Echo {
	app := echo.New()
	conf := database.Config{}

	routes.RootRoutes(app, conf)
	app.Use(middleware.CORS())
	app.Use(middleware.Logger())

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": true,
		})
	})

	return app
}
