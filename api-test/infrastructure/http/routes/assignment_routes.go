package routes

import (
	"api-test/controller"
	"api-test/infrastructure/database"
	"api-test/infrastructure/http/middleware"
	"api-test/repository"
	"api-test/service"

	"github.com/labstack/echo/v4"
)

func RoutesAssignment(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	assignmentRepository := repository.NewAssignmentRepository(db)
	assignmentService := service.NewAssignmentService(assignmentRepository, conf)
	assignmentController := controller.NewAssignmentController(assignmentService)

	teacher := echo.Group("/v1/teacher")

	teacher.POST("/assignment", assignmentController.CreateAssignment, middleware.JWTTokenMiddleware())
	echo.GET("/assignment", assignmentController.GetAllAssignment)
	echo.GET("/assignment/:id", assignmentController.GetAssignmentById)
	teacher.PUT("/assignment/:id", assignmentController.UpdateAssignment, middleware.JWTTokenMiddleware())
	teacher.DELETE("/assignment/:id", assignmentController.DeleteAssignment)
}
