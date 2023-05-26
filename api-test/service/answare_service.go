package service

import (
	"api-test/infrastructure/database"
	"api-test/interfaces"
	"api-test/models"
)

type svcAnsware struct {
	c    database.Config
	repo interfaces.AnswareRepository
}

func NewAnswareService(repo interfaces.AnswareRepository, c database.Config) interfaces.AnswareService {
	return &svcAnsware{
		repo: repo,
		c:    c,
	}
}

func (s *svcAnsware) CreateAnswareService(answare models.Answare) (string, error) {
	if answare.UserId == 1 {
		return "User Can't Answare Because The User is Teacher", nil
	} else {
		return "Student Create Answare From Question ...", s.repo.CreateAnsware(answare)
	}
}

func (s *svcAnsware) GetAnswareByIdService(id int) (models.Answare, error) {
	var answare models.Answare
	if id != answare.QuestionId {
		return answare, nil
	} else {
		return s.repo.GetAnswareById(id)
	}
}

func (s *svcAnsware) GetAllAnswareService() (ans []models.Answare, err error) {
	return s.repo.GetAllAnsware()
}

func (s *svcAnsware) UpdateAnswareService(id int, answare models.Answare) (models.Answare, error) {
	if answare.UserId == 1 && answare.QuestionId == id {
		return answare, nil
	} else {
		return s.repo.UpdateAnsware(id, answare)
	}
}

func (s *svcAnsware) DeleteAnswareService(id int, answare models.Answare) (string, error) {
	if answare.UserId == 1 && answare.QuestionId == id {
		return "user can't delete the answare because user is teacher", nil
	} else {
		return "delete answare from question", s.repo.DeleteAnsware(id, answare)
	}
}
