package input

import "go-hexagonal/internal/domain"

type UserServicePort interface {
	CreateUser(user domain.User) error
	GetUserByID(id string) (*domain.User, error)
	UpdateUser(user domain.User) error
	DeleteUser(id string) error
}
