package repository

import (
	"api-test/interfaces"
	"api-test/models"

	"gorm.io/gorm"
)

type discussionsRepository struct {
	DB *gorm.DB
}

func NewDiscussionsRepository(db *gorm.DB) interfaces.DiscussionsRepository {
	return &discussionsRepository{
		DB: db,
	}
}

func (r *discussionsRepository) CreateDiscussionsRepository(discussions models.Discussions) error {
	var user models.User
	var questions models.Question
	var answare models.Answare
	r.DB.Where("id = ?", discussions.QuestionId).Find(&questions)
	r.DB.Where("id = ? ", discussions.UserId).Find(&user)
	r.DB.Where("id = ?", discussions.AnswareId).Find(&answare)
	discussions.Name = user.Name
	discussions.QuestionUser = questions.QuestionUser
	discussions.AnswereUser = answare.AnswareUser

	r.DB.Create(&discussions)
	return nil
}

func (r *discussionsRepository) GetAllDiscussionsRepository() (discussions []models.Discussions, err error) {
	r.DB.Find(&discussions)
	return discussions, nil
}
