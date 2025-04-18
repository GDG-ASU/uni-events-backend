package repositories

import (
	"context"
	"uni-events-backend/internal/models"

	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error) {
	if err := r.db.WithContext(ctx).Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}
