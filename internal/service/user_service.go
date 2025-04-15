package service

import (
	"context"
	"uni-events-backend/internal/models"
	"uni-events-backend/internal/repositories"
)

type UserService interface {
	GetUserByClerkID(ctx context.Context, clerkID string) (*models.User, error)
	CreateUserIfNotExists(ctx context.Context, user *models.User) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetUserByClerkID(ctx context.Context, clerkID string) (*models.User, error) {
	return s.userRepo.FindByClerkID(ctx, clerkID)
}

func (s *userService) CreateUserIfNotExists(ctx context.Context, user *models.User) (*models.User, error) {
	existing, err := s.userRepo.FindByClerkID(ctx, user.ClerkID)
	if err == nil && existing != nil {
		return existing, nil
	}
	return s.userRepo.Create(ctx, user)
}
