package service

import (
	"dev11/model"
	"dev11/repo"
	"time"
)

type EventService interface {
	Create(user_id uint64, e model.Event) (*model.Event, error)
	Update(user_id uint64, e model.Event) error
	Delete(user_id uint64, event_id uint64) error
	AllForDay(user_id uint64, day time.Time) ([]model.Event, error)
	AllForWeek(user_id uint64, day time.Time) ([]model.Event, error)
	AllForMonth(user_id uint64, day time.Time) ([]model.Event, error)
}

type EventServiceImpl struct {
	repo repo.EventRepository
}

// * Create
func (s *EventServiceImpl) Create(user_id uint64, e model.Event) (*model.Event, error) {
	event, err := s.repo.Create(user_id, e)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// * Update
func (s *EventServiceImpl) Update(user_id uint64, event model.Event) error {
	err := s.repo.Update(user_id, event)
	if err != nil {
		return err
	}
	return nil
}

// * Delete
func (s *EventServiceImpl) Delete(user_id uint64, event_id uint64) error {
	err := s.repo.Delete(user_id, event_id)
	if err != nil {
		return err
	}
	return nil
}

// * All for a day
func (s *EventServiceImpl) AllForDay(user_id uint64, day time.Time) ([]model.Event, error) {
	events, err := s.repo.GetForDay(user_id, day)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// * All for a week
func (s *EventServiceImpl) AllForWeek(user_id uint64, day time.Time) ([]model.Event, error) {
	events, err := s.repo.GetForWeek(user_id, day)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// * All for a month
func (s *EventServiceImpl) AllForMonth(user_id uint64, day time.Time) ([]model.Event, error) {
	events, err := s.repo.GetForMonth(user_id, day)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func NewUserService(repo repo.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{repo: repo}
}
