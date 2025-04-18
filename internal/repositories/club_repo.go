package repositories

import (
	"uni-events-backend/internal/models"

	"gorm.io/gorm"
)

type ClubRepository interface {
	CreateClub(club *models.Club) error
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