package repository

import (
	"api-test/interfaces"
	"api-test/models"

	"gorm.io/gorm"
)

type assignmentRepository struct {
	DB *gorm.DB
}

func NewAssignmentRepository(db *gorm.DB) interfaces.AssignmentRepository {
	return &assignmentRepository{
		DB: db,
	}
}

func (r *assignmentRepository) CreateAssignmentRepository(assignment models.Assignment) error {
	var user models.User
	r.DB.Where("id = ?", assignment.UserId).First(&user)
	assignment.Name = user.Name
	r.DB.Create(&assignment)
	return nil
}

func (r *assignmentRepository) GetAssignmentByIdRepository(id int) (models.Assignment, error) {
	var assign models.Assignment
	var user models.User
	r.DB.Joins("JOIN users ON users.id = assignments.user_id").Where("assignments.id = ?", id).First(&assign)
	r.DB.Where("id = ?", assign.UserId).First(&user)
	assign.Name = user.Name
	return assign, nil
}

func (r *assignmentRepository) GetAllAssignmentRepository() ([]models.Assignment, error) {
	assignments := []models.Assignment{}
	r.DB.Find(&assignments)
	return assignments, nil
}

func (r *assignmentRepository) UpdateAssignmentRepository(id int, assignment models.Assignment) (models.Assignment, error) {
	var user models.User
	r.DB.Model(&assignment).Where("id = ?", id).Updates(assignment)
	r.DB.Joins("JOIN users ON users.id = assignments.user_id").Where("assignments.id = ?", id).First(&assignment)
	r.DB.Where("id = ?", assignment.UserId).First(&user)
	assignment.Name = user.Name
	return assignment, nil
}

func (r *assignmentRepository) DeleteAssignmentRepository(id int, assign models.Assignment) error {
	r.DB.Where("id = ?", id).Delete(&assign)
	return nil
}
