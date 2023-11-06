package output

import "go-hexagonal/internal/domain"

type UserRepositoryPort interface {
	Save(domain.User) error
	FindByID(string) (*domain.User, error)
	Update(domain.User) error
}
