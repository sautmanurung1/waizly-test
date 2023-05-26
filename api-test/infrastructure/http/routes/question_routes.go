package routes

import (
	"api-test/controller"
	"api-test/infrastructure/database"
	"api-test/infrastructure/http/middleware"
	"api-test/repository"
	"api-test/service"

	"github.com/labstack/echo/v4"
)

func RoutesQuestion(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	questionRepository := repository.NewQuestionRepository(db)
	questionService := service.NewQuestionService(questionRepository, conf)
	questionController := controller.NewQuestionController(questionService)

	teacher := echo.Group("/teacher")
	teacher.POST("/questions", questionController.CreateQuestion, middleware.JWTTokenMiddleware())
	echo.GET("/questions", questionController.GetAllQuestion)
	echo.GET("/questions/:id", questionController.GetQuestionById)
	teacher.PUT("/questions/:id", questionController.UpdateQuestion, middleware.JWTTokenMiddleware())
	teacher.DELETE("/questions/:id", questionController.DeleteQuestion, middleware.JWTTokenMiddleware())
}
