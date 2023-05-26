package service

import (
	"api-test/infrastructure/database"
	"api-test/infrastructure/http/middleware"
	"api-test/interfaces"
	"api-test/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type svcAuth struct {
	c    database.Config
	repo interfaces.AuthRepository
}

func NewAuthService(repo interfaces.AuthRepository, c database.Config) interfaces.AuthService {
	return &svcAuth{
		repo: repo,
		c:    c,
	}
}

func (s *svcAuth) LoginService(username string, password string, roleId int) (string, int) {
	user, _ := s.repo.LoginRepository(username)

	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if er != nil {
		return "Wrong Password", http.StatusUnauthorized
	}

	if user.RoleId != roleId {
		return "Your Role Id error", http.StatusUnauthorized
	}

	token, _ := middleware.CreateToken(int(user.ID), user.Username, s.c.JWT_KEY)

	return token, http.StatusOK
}

func (s *svcAuth) RegisterService(creadential models.User) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(creadential.Password), bcrypt.MinCost)
	creadential.Password = string(password)
	return s.repo.RegisterRepository(creadential)
}
