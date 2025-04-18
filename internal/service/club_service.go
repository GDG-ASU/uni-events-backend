package service

import (
	"context"
	"uni-events-backend/internal/models"
	"uni-events-backend/internal/repositories"
)

type ClubService interface {
	CreateClub(ctx context.Context, club *models.Club) (*models.Club, error)
	UpdateClub(ctx context.Context, clubID uint, name string, description string) (*models.Club, error)
	IsUserClubOwner(ctx context.Context, clubID, userID uint) (bool, error)
	GetClubByID(ctx context.Context, clubID uint) (*models.Club, error)
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

func (s *clubService) GetClubByID(ctx context.Context, clubID uint) (*models.Club, error) {
	return s.repo.GetClubByID(ctx, clubID)
}

func (s *clubService) UpdateClub(ctx context.Context, clubID uint, name, description string) (*models.Club, error) {
	club, err := s.repo.GetClubByID(ctx, clubID)
	if err != nil {
		return nil, err
	}

	club.Name = name
	club.Description = description

	if err := s.repo.UpdateClub(ctx, club); err != nil {
		return nil, err
	}

	return club, nil
}

func (s *clubService) IsUserClubOwner(ctx context.Context, clubID, userID uint) (bool, error) {
	return s.repo.IsUserClubOwner(ctx, clubID, userID)
}