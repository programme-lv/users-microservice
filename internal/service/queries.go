package service

import (
	"github.com/google/uuid"
	"github.com/programme-lv/users-microservice/internal/domain"
)

func (s *UserService) GetUser(uuid uuid.UUID) (domain.User, error) {
	return s.repo.GetUser(uuid)
}

func (s *UserService) ListUsers() ([]domain.User, error) {
	return s.repo.ListUsers()
}
