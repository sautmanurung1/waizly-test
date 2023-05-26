package routes

import (
	"api-test/controller"
	"api-test/infrastructure/database"
	"api-test/infrastructure/http/middleware"
	"api-test/repository"
	"api-test/service"

	"github.com/labstack/echo/v4"
)

func RoutesAnsware(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	answareRepository := repository.NewAnswareRepository(db)
	answareService := service.NewAnswareService(answareRepository, conf)
	answareController := controller.NewAnswareController(answareService)

	answare := echo.Group("/v1/student")

	answare.POST("/answare", answareController.CreateAnsware, middleware.JWTTokenMiddleware())
	answare.GET("/answare", answareController.GetAllAnsware)
	echo.GET("/answare/:id", answareController.GetAnswareById)
	echo.PUT("/answare/:id", answareController.UpdateAnsware, middleware.JWTTokenMiddleware())
	echo.DELETE("/answare/:id", answareController.DeleteAnsware, middleware.JWTTokenMiddleware())
}
