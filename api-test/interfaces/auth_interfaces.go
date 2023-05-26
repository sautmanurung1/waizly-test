package interfaces

import (
	"api-test/models"

	"github.com/labstack/echo/v4"
)

type AuthRepository interface {
	LoginRepository(username string) (user models.User, err error)
	RegisterRepository(user models.User) error
}

type AuthService interface {
	LoginService(username string, password string, roleId int) (string, int)
	RegisterService(creadential models.User) error
}

type AuthHandler interface {
	LoginUser(c echo.Context) error
	RegisterUser(c echo.Context) error
}
