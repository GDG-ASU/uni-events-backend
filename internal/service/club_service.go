package service

import (
	"context"
	"uni-events-backend/internal/models"
	"uni-events-backend/internal/repositories"
)

type ClubService interface {
	CreateClub(ctx context.Context, club *models.Club) (*models.Club, error)
}

type clubService struct {
	repo repositories.ClubRepository
}

func NewClubService(repo repositories.ClubRepository) ClubService {
	return &clubService{repo}
}

func (s *clubService) CreateClub(ctx context.Context, club *models.Club) (*models.Club, error) {
	err := s.repo.CreateClub(club)
	return club, err
}
