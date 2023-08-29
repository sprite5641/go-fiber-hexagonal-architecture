package user

import (
	"go-hexagonal/internal/domain"
	"go-hexagonal/internal/ports/input"
)

type UserService struct {
	repo domain.UserRepository
}

// Ensure UserService implements UserServicePort
var _ input.UserServicePort = &UserService{}

func NewUserService(r domain.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(user domain.User) error {
	return s.repo.Save(user)
}

func (s *UserService) GetUserByID(id string) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) UpdateUser(user domain.User) error {
	// Implement your update logic here
	return nil
}

func (s *UserService) DeleteUser(id string) error {
	// Implement your delete logic here
	return nil
}
