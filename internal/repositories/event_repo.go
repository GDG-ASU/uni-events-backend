package repositories

import (
	"context"
	"uni-events-backend/internal/models"

	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, event *models.Event) (*models.Event, error)
	GetAllEvents(ctx context.Context) ([]*models.Event, error)
	UpdateEvent(ctx context.Context, event *models.Event) (*models.Event, error)
	DeleteEvent(ctx context.Context, id *models.Event) (*models.Event, error)
	GetEventByID(ctx context.Context, id uint) (*models.Event, error)
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

func (r *eventRepository) GetAllEvents(ctx context.Context) ([]*models.Event, error) {
	var events []*models.Event
	if err := r.db.WithContext(ctx).Preload("Club").Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepository) GetEventByID(ctx context.Context, id uint) (*models.Event, error) {
	var event models.Event
	if err := r.db.WithContext(ctx).First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) UpdateEvent(ctx context.Context, event *models.Event) (*models.Event, error) {
	if err := r.db.WithContext(ctx).Save(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *eventRepository) DeleteEvent(ctx context.Context, event *models.Event) (*models.Event, error) {
	if err := r.db.WithContext(ctx).Delete(event).Error; err != nil {
		return nil,err
	}
	return event,nil
}
