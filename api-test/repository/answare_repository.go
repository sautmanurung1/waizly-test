package repository

import (
	"api-test/interfaces"
	"api-test/models"

	"gorm.io/gorm"
)

type repositoryAnsware struct {
	DB *gorm.DB
}

func NewAnswareRepository(db *gorm.DB) interfaces.AnswareRepository {
	return &repositoryAnsware{
		DB: db,
	}
}

func (r *repositoryAnsware) CreateAnsware(answare models.Answare) error {
	var questions models.Question
	r.DB.Where("id = ?", answare.QuestionId).Preload("Question").Find(&questions)
	answare.Questions = questions.QuestionUser
	answare.Name = questions.Name
	r.DB.Create(&answare)
	return nil
}

func (r *repositoryAnsware) GetAnswareById(id int) (models.Answare, error) {
	var ans models.Answare
	var questions models.Question
	r.DB.Joins("JOIN questions ON questions.id = answares.question_id").Where("answares.id = ?", id).First(&ans)
	r.DB.Where("id = ?", ans.QuestionId).Preload("Question").Find(&questions)
	ans.Name = questions.Name
	ans.Questions = questions.QuestionUser
	return ans, nil
}

func (r *repositoryAnsware) GetAllAnsware() (ans []models.Answare, err error) {
	r.DB.Find(&ans)
	return ans, nil
}

func (r *repositoryAnsware) UpdateAnsware(id int, answare models.Answare) (models.Answare, error) {
	var questions models.Question
	r.DB.Model(&answare).Where("id = ?", id).Updates(&answare)
	r.DB.Joins("JOIN questions ON questions.id = answares.question_id").Where("answares.id = ?", id).First(&answare)
	r.DB.Where("id = ?", answare.QuestionId).First(&questions)
	answare.Name = questions.Name
	answare.Questions = questions.QuestionUser
	return answare, nil
}

func (r *repositoryAnsware) DeleteAnsware(id int, ans models.Answare) error {
	r.DB.Where("id = ?", id).Delete(&ans)
	return nil
}
