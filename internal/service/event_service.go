package service

import (
	"context"
	"uni-events-backend/internal/models"
	"uni-events-backend/internal/repositories"
)

type EventService interface {
	CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error)
	GetAllEvents(ctx context.Context) ([]*models.Event, error)
	UpdateEvent(ctx context.Context, event *models.Event) (*models.Event, error)
	DeleteEvent(ctx context.Context, event *models.Event) (*models.Event, error)
	GetEventByID(ctx context.Context, id uint) (*models.Event, error)
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

func (s *eventService) GetAllEvents(ctx context.Context) ([]*models.Event, error) {
	return s.repo.GetAllEvents(ctx)
}

func (s *eventService) GetEventByID(ctx context.Context, id uint) (*models.Event, error) {
	return s.repo.GetEventByID(ctx, id)
}

func (s *eventService) UpdateEvent(ctx context.Context, event *models.Event) (*models.Event, error) {
	return s.repo.UpdateEvent(ctx, event)
}

func (s *eventService) DeleteEvent(ctx context.Context, event *models.Event) (*models.Event, error) {
	return s.repo.DeleteEvent(ctx, event)
}