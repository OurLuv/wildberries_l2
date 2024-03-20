package model

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Event struct {
	ID    uint64    `json:"id,omitempty"`
	Title string    `json:"title,omitempty"`
	Date  time.Time `json:"date,omitempty"`
}

func (e *Event) Serialize() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Event) ParseAndValidate(data url.Values) error {
	title := data.Get("title")
	if title == "" {
		return fmt.Errorf("missing Title parameter")
	}
	e.Title = title

	dateStr := data.Get("date")
	if dateStr == "" {
		return fmt.Errorf("missing Date parameter")
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("invalid Date parameter: %v", err)
	}
	e.Date = date

	return nil
}
