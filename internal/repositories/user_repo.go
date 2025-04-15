package repositories

import (
	"context"
	"uni-events-backend/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByClerkID(ctx context.Context, clerkID string) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) FindByClerkID(ctx context.Context, clerkID string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("clerk_id = ?", clerkID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Create(ctx context.Context, user *models.User) (*models.User, error) {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
