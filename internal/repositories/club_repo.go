package repositories

import (
	"context"
	"uni-events-backend/internal/models"

	"gorm.io/gorm"
)

type ClubRepository interface {
	CreateClub(club *models.Club) error
	UpdateClub(ctx context.Context, club *models.Club) error
	GetClubByID(ctx context.Context, clubID uint) (*models.Club, error)
	IsUserClubOwner(ctx context.Context, clubID uint, userID uint) (bool, error)
}

type clubRepo struct {
	db *gorm.DB
}

func NewClubRepository(db *gorm.DB) ClubRepository {
	return &clubRepo{db}
}

func (r *clubRepo) CreateClub(club *models.Club) error {
	return r.db.Create(club).Error
}

func (r *clubRepo) GetClubByID(ctx context.Context, id uint) (*models.Club, error) {
	var club models.Club
	err := r.db.WithContext(ctx).First(&club, id).Error
	return &club, err
}

func (r *clubRepo) UpdateClub(ctx context.Context, club *models.Club) error {
	return r.db.WithContext(ctx).Save(club).Error
}

func (r *clubRepo) IsUserClubOwner(ctx context.Context, clubID uint, userID uint) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.ClubOwner{}).
		Where("club_id = ? AND user_id = ?", clubID, userID).Count(&count).Error
	return count > 0, err
}