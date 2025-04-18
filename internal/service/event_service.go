package service

import (
	"context"
	"uni-events-backend/internal/models"
	"uni-events-backend/internal/repositories"
)

type EventService interface {
	CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error)
}

type eventService struct {
	repo repositories.EventRepository
}

func NewEventService(repo repositories.EventRepository) EventService {
	return &eventService{repo}
}

func (s *eventService) CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error) {
	return s.repo.CreateEvent(ctx, event)
}
