package routes

import (
	"api-test/controller"
	"api-test/infrastructure/database"
	"api-test/repository"
	"api-test/service"

	"github.com/labstack/echo/v4"
)

func RoutesAuth(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, conf)
	authController := controller.NewAuthController(authService)

	auth := echo.Group("/v1/auth")

	auth.POST("/register", authController.RegisterUser)
	auth.POST("/login", authController.LoginUser)
}
