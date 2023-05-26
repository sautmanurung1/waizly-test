package controller

import (
	"api-test/interfaces"
	"api-test/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authController struct {
	s interfaces.AuthService
}

func NewAuthController(authService interfaces.AuthService) interfaces.AuthHandler {
	return &authController{
		s: authService,
	}
}

func (a *authController) LoginUser(c echo.Context) error {
	userLogin := models.User{}

	err := c.Bind(&userLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Bad Request",
			"error":  err.Error(),
		})
	}

	token, statusCode := a.s.LoginService(userLogin.Username, userLogin.Password, userLogin.RoleId)

	if statusCode == http.StatusUnauthorized {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":    "your username and password wrong",
			"statusCode": statusCode,
		})
	} else if statusCode == http.StatusInternalServerError {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":    "internal server error",
			"statusCode": statusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"data":    token,
	})
}

func (a *authController) RegisterUser(c echo.Context) error {
	user := models.User{}
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	er := a.s.RegisterService(user)

	if er != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "user UnRegistered",
			"data":    er.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success register",
		"data":    user,
	})
}
