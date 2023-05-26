package service

import (
	"api-test/infrastructure/database"
	"api-test/interfaces"
	"api-test/models"
)

type svcQuestion struct {
	c    database.Config
	repo interfaces.QuestionRepository
}

func NewQuestionService(repo interfaces.QuestionRepository, c database.Config) interfaces.QuestionService {
	return &svcQuestion{
		c:    c,
		repo: repo,
	}
}

func (s *svcQuestion) CreateQuestionService(question models.Question) (string, error) {
	if question.UserId == 1 {
		return "teacher create question from assignment", s.repo.CreateQuestionRepository(question)
	} else {
		return "user can't make question because the user is student", nil
	}
}

func (s *svcQuestion) GetQuestionByIdService(id int) (models.Question, error) {
	return s.repo.GetQuestionByIdRepository(id)
}

func (s *svcQuestion) GetAllQuestionService() (question []models.Question, err error) {
	return s.repo.GetAllQuestionRepository()
}

func (s *svcQuestion) UpdateQuestionService(id int, question models.Question) (models.Question, error) {
	if question.UserId == 1 && question.AssignmentId == id {
		return s.repo.UpdateQuestionRepository(id, question)
	} else {
		return question, nil
	}
}

func (s *svcQuestion) DeleteQuestionService(id int, question models.Question) (string, error) {
	if question.UserId == 1 && question.AssignmentId == id {
		return "delete question by id", s.repo.DeleteQuestionRepository(id, question)
	} else {
		return "user can't update question because user is student", nil
	}
}
