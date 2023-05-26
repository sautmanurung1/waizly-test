package routes

import (
	"api-test/controller"
	"api-test/infrastructure/database"
	"api-test/repository"
	"api-test/service"

	"github.com/labstack/echo/v4"
)

func RoutesDiscussions(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	discussionsRepository := repository.NewDiscussionsRepository(db)
	discussiosService := service.NewDiscussionService(discussionsRepository, conf)
	discussionsController := controller.NewDiscussionsController(discussiosService)

	echo.GET("/discussions", discussionsController.CreateDiscussions)
	echo.POST("/discussions", discussionsController.CreateDiscussions)
}
