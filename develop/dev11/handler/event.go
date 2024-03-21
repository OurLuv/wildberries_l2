package handler

import (
	"dev11/model"
	"net/http"
	"strconv"
	"time"
)

// POST /create_event
// POST /update_event
// POST /delete_event
// GET /events_for_day
// GET /events_for_week
// GET /events_for_month

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			SendError(w, "failed to parse form data: "+err.Error(), http.StatusBadRequest)
			return
		}

		// event - parsing & validate
		event := model.Event{}
		if err := event.ParseAndValidate(r.Form); err != nil {
			SendError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// user - parsing & validate
		user := model.User{}
		if err := user.ParseAndValidate(r.Form); err != nil {
			SendError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// service
		result, err := h.EventService.Create(user.ID, event)
		if err != nil {
			SendError(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		SendResult(w, *result)

	} else {
		SendError(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			SendError(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
			return
		}

		// event - parsing & validate
		event := model.Event{}
		if err := event.ParseAndValidate(r.Form); err != nil {
			SendError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// user - parsing & validate
		user := model.User{}
		if err := user.ParseAndValidate(r.Form); err != nil {
			SendError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// service
		if err := h.EventService.Update(user.ID, event); err != nil {
			SendError(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		SendResult(w, "event is updated")

	} else {
		SendError(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// * Delete
func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			SendError(w, "failed to parse form data: "+err.Error(), http.StatusBadRequest)
			return
		}

		event_id, err := GetIdFromReq(r.Form, "event_id")
		if err != nil {
			SendError(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
			return
		}

		user_id, err := GetIdFromReq(r.Form, "user_id")
		if err != nil {
			SendError(w, "failed to parse form data: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.EventService.Delete(user_id, event_id); err != nil {
			SendError(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		SendResult(w, "event is deleted")

	} else {
		SendError(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// * Get events for a day
func (h *Handler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		user_idStr := r.URL.Query().Get("user_id")
		if user_idStr == "" {
			SendError(w, "missing user id", http.StatusBadRequest)
			return
		}
		user_idInt, err := strconv.Atoi(user_idStr)
		user_id := uint64(user_idInt)
		if err != nil {
			SendError(w, "can't parse a user id", http.StatusBadRequest)
			return
		}
		dayStr := r.URL.Query().Get("date")
		if dayStr == "" {
			SendError(w, "missing day", http.StatusBadRequest)
			return
		}
		day, err := time.Parse("2006-01-02", dayStr)
		if err != nil {
			SendError(w, "can't parse a date", http.StatusBadRequest)
			return
		}

		events, err := h.EventService.AllForDay(user_id, day)
		if err != nil {
			SendError(w, "can't parse a date", http.StatusServiceUnavailable)
			return
		}

		SendResult(w, events)
	} else {
		SendError(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// * Get events for a week
func (h *Handler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		user_idStr := r.URL.Query().Get("user_id")
		if user_idStr == "" {
			SendError(w, "missing user id", http.StatusBadRequest)
			return
		}
		user_idInt, err := strconv.Atoi(user_idStr)
		user_id := uint64(user_idInt)
		if err != nil {
			SendError(w, "can't parse a user id", http.StatusBadRequest)
			return
		}
		dayStr := r.URL.Query().Get("date")
		if dayStr == "" {
			SendError(w, "missing day", http.StatusBadRequest)
			return
		}
		day, err := time.Parse("2006-01-02", dayStr)
		if err != nil {
			SendError(w, "can't parse a date", http.StatusBadRequest)
			return
		}

		events, err := h.EventService.AllForWeek(user_id, day)
		if err != nil {
			SendError(w, "can't parse a date", http.StatusServiceUnavailable)
			return
		}

		SendResult(w, events)
	} else {
		SendError(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// * Get events for a month
func (h *Handler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		user_idStr := r.URL.Query().Get("user_id")
		if user_idStr == "" {
			SendError(w, "missing user id", http.StatusBadRequest)
			return
		}
		user_idInt, err := strconv.Atoi(user_idStr)
		user_id := uint64(user_idInt)
		if err != nil {
			SendError(w, "can't parse a user id", http.StatusBadRequest)
			return
		}
		dayStr := r.URL.Query().Get("date")
		if dayStr == "" {
			SendError(w, "missing day", http.StatusBadRequest)
			return
		}
		day, err := time.Parse("2006-01-02", dayStr)
		if err != nil {
			SendError(w, "can't parse a date", http.StatusBadRequest)
			return
		}

		events, err := h.EventService.AllForMonth(user_id, day)
		if err != nil {
			SendError(w, "can't parse a date", http.StatusServiceUnavailable)
			return
		}

		SendResult(w, events)
	} else {
		SendError(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
