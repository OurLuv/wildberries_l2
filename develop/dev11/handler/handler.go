package handler

import (
	"dev11/service"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Handler struct {
	service.EventService
}

type Response struct {
	Result any    `json:"result,omitempty"`
	Err    string `json:"error,omitempty"`
}

func (h *Handler) Init() error {
	// event
	http.HandleFunc("/create_event", loggingRequest(h.CreateEvent))
	http.HandleFunc("/update_event", loggingRequest(h.UpdateEvent))
	http.HandleFunc("/delete_event", loggingRequest(h.DeleteEvent))
	http.HandleFunc("/events_for_day", loggingRequest(h.GetEventsForDay))
	http.HandleFunc("/events_for_week", loggingRequest(h.GetEventsForWeek))
	http.HandleFunc("/events_for_month", loggingRequest(h.GetEventsForMonth))

	return http.ListenAndServe(":8080", nil)
}

func SendError(w http.ResponseWriter, errorStr string, code int) {
	w.WriteHeader(code)
	resp := Response{
		Err: errorStr,
	}
	json.NewEncoder(w).Encode(resp)
}

func SendResult(w http.ResponseWriter, data any) {
	resp := Response{
		Result: data,
	}
	json.NewEncoder(w).Encode(resp)
}

func GetIdFromReq(data url.Values, field string) (uint64, error) {
	idStr := data.Get(field)
	if idStr == "" {
		return 0, fmt.Errorf("missing ID parameter")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid ID parameter: %v", err)
	}
	return uint64(id), nil
}

func NewHandler(s service.EventService) *Handler {
	return &Handler{
		EventService: s,
	}
}
