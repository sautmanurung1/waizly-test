package service

import (
	"api-test/infrastructure/database"
	"api-test/interfaces"
	"api-test/models"
)

type svcAssignment struct {
	c    database.Config
	repo interfaces.AssignmentRepository
}

func NewAssignmentService(repo interfaces.AssignmentRepository, c database.Config) interfaces.AssignmentService {
	return &svcAssignment{
		c:    c,
		repo: repo,
	}
}

func (s *svcAssignment) CreateAssignmentService(assignment models.Assignment) (string, error) {
	if assignment.UserId != 1 {
		return "you're not teacher", nil
	} else {
		return "teacher can create assignment", s.repo.CreateAssignmentRepository(assignment)
	}
}

func (s *svcAssignment) GetAssginmentByIdService(id int) (models.Assignment, error) {
	return s.repo.GetAssignmentByIdRepository(id)
}

func (s *svcAssignment) GetAllAssignmentService() (assignment []models.Assignment, err error) {
	return s.repo.GetAllAssignmentRepository()
}

func (s *svcAssignment) UpdateAssignmentService(id int, assignment models.Assignment) (models.Assignment, error) {
	if assignment.UserId != 1 {
		return assignment, nil
	} else {
		return s.repo.UpdateAssignmentRepository(id, assignment)
	}
}

func (s *svcAssignment) DeleteAssignmentService(id int, assignment models.Assignment) (string, error) {
	if assignment.UserId != 1 {
		return "you can't delete the assignment", nil
	} else {
		return "you can delete the assignment by id", s.repo.DeleteAssignmentRepository(id, assignment)
	}
}
