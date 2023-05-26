package service

import (
	"api-test/infrastructure/database"
	"api-test/interfaces"
	"api-test/models"
)

type svcDiscussions struct {
	c    database.Config
	repo interfaces.DiscussionsRepository
}

func NewDiscussionService(repo interfaces.DiscussionsRepository, c database.Config) interfaces.DiscussionsService {
	return &svcDiscussions{
		c:    c,
		repo: repo,
	}
}

func (s *svcDiscussions) CreateDiscussionsService(discussions models.Discussions) error {
	return s.repo.CreateDiscussionsRepository(discussions)
}

func (s *svcDiscussions) GetAllDiscussionsService() (discussions []models.Discussions, err error) {
	return s.repo.GetAllDiscussionsRepository()
}
