package user

import (
	"context"
	"go-hexagonal/internal/domain"
	"go-hexagonal/internal/infrastructure/repository/redis"
	"go-hexagonal/internal/ports/input"
)

// Ensure UserService implements UserServicePort
var _ input.UserServicePort = &UserService{}

type UserService struct {
	repo        domain.UserRepository
	redisClient *redis.RedisClient
}

func NewUserService(r domain.UserRepository, redis *redis.RedisClient) *UserService {
	return &UserService{
		repo:        r,
		redisClient: redis,
	}
}

func (s *UserService) CreateUser(user domain.User) error {
	return s.repo.Save(user)
}

func (s *UserService) GetUserByID(id string) (*domain.User, error) {

	ctx := context.Background()

	username, err := s.redisClient.GetKey(ctx, id)

	if err != nil || username == "" {
		user, err := s.repo.FindByID(id)
		if err != nil {
			return nil, err
		}
		err = s.redisClient.SetKey(ctx, id, user.Username)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	return &domain.User{Username: username}, nil
}

func (s *UserService) UpdateUser(user domain.User) error {
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
