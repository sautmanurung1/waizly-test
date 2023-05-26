package repository

import (
	"api-test/interfaces"
	"api-test/models"
	"database/sql"

	"gorm.io/gorm"
)

type questionRepository struct {
	DB *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) interfaces.QuestionRepository {
	return &questionRepository{
		DB: db,
	}
}

func (r *questionRepository) CreateQuestionRepository(question models.Question) error {
	var assignments models.Assignment
	var users models.User
	r.DB.Where("id = ?", question.AssignmentId).First(&assignments)
	r.DB.Where("id = ?", question.UserId).First(&users)
	r.DB.Raw("SELECT title FROM assignments WHERE title = @title",
		sql.Named("title", assignments.Title)).Find(&assignments)
	question.Name = assignments.Name
	question.AssignmentTitle = assignments.Title
	r.DB.Create(&question)
	return nil
}

func (r *questionRepository) GetQuestionByIdRepository(id int) (models.Question, error) {
	var questions models.Question
	var assign models.Assignment
	r.DB.Joins("JOIN assignments ON assignments.id = questions.assignment_id").Where("questions.id = ?", id).First(&questions)
	r.DB.Where("id = ?", questions.AssignmentId).First(&assign)
	questions.Name = assign.Name
	return questions, nil
}

func (r *questionRepository) GetAllQuestionRepository() ([]models.Question, error) {
	questions := []models.Question{}
	r.DB.Find(&questions)
	return questions, nil
}

func (r *questionRepository) UpdateQuestionRepository(id int, question models.Question) (models.Question, error) {
	var assignments models.Assignment
	r.DB.Model(&question).Where("id = ?", id).Updates(&question)
	r.DB.Joins("JOIN assignments ON assignments.id = questions.assignment_id").Where("questions.id = ?", id).First(&question)
	r.DB.Where("id = ?", question.AssignmentId).First(&assignments)
	question.Name = assignments.Name
	return question, nil
}

func (r *questionRepository) DeleteQuestionRepository(id int, question models.Question) error {
	r.DB.Where("id = ?", id).Delete(&question)
	return nil
}
