package repo

import (
	"dev11/model"
	"time"
)

type EventRepository interface {
	Create(user_id uint64, e model.Event) (model.Event, error)
	Update(user_id uint64, e model.Event) error
	Delete(user_id uint64, event_id uint64) error
	GetForDay(user_id uint64, day time.Time) ([]model.Event, error)
	GetForWeek(user_id uint64, week time.Time) ([]model.Event, error)
	GetForMonth(user_id uint64, month time.Time) ([]model.Event, error)
}

type Mock struct{}

func (mr *Mock) Create(user_id uint64, e model.Event) (model.Event, error) {
	event := model.Event{
		ID:    12,
		Title: "Meeting",
		Date:  time.Time{},
	}
	return event, nil
}
func (mr *Mock) Update(user_id uint64, e model.Event) error   { return nil }
func (mr *Mock) Delete(user_id uint64, event_id uint64) error { return nil }
func (mr *Mock) GetForDay(user_id uint64, day time.Time) ([]model.Event, error) {
	events := []model.Event{
		{
			ID:    30,
			Title: "Meet up",
			Date:  time.Time{},
		},
		{
			ID:    33,
			Title: "Interview with AC3-group",
			Date:  time.Time{},
		},
	}
	return events, nil
}
func (mr *Mock) GetForWeek(user_id uint64, week time.Time) ([]model.Event, error) {
	events := []model.Event{
		{
			ID:    30,
			Title: "Meet up",
			Date:  time.Time{},
		},
		{
			ID:    33,
			Title: "Interview with AC3-group",
			Date:  time.Time{},
		},
		{
			ID:    34,
			Title: "Lecture in class 201b",
			Date:  time.Time{},
		},
	}
	return events, nil
}
func (mr *Mock) GetForMonth(user_id uint64, month time.Time) ([]model.Event, error) {
	events := []model.Event{
		{
			ID:    30,
			Title: "Meet up",
			Date:  time.Time{},
		},
		{
			ID:    33,
			Title: "Interview with AC3-group",
			Date:  time.Time{},
		},
		{
			ID:    34,
			Title: "Lecture in class 201b",
			Date:  time.Time{},
		},
		{
			ID:    37,
			Title: "Book the hotel in Warsaw",
			Date:  time.Time{},
		},
	}
	return events, nil
}
