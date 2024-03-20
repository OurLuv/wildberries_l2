package main

import (
	"dev11/handler"
	"dev11/repo"
	"dev11/service"
	"log"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var h handler.Handler

func TestMain(m *testing.M) {
	repo := repo.Mock{}
	service := service.NewUserService(&repo)
	h = *handler.NewHandler(service)
	go func() {
		if err := h.Init(); err != nil {
			log.Fatal(err)
		}
	}()

	os.Exit(m.Run())
}

func TestCreateEvent(t *testing.T) {
	testCases := []struct {
		name               string
		inputBody          string
		ExpectedStatusCode int
	}{
		{
			name:               "status ok",
			inputBody:          "user_id=3&event_id=12&title=Meeting&date=2019-09-09",
			ExpectedStatusCode: 200,
		},
		{
			name:               "not correct user id",
			inputBody:          "user_id=3d&title=Meeting&date=2019-09-09",
			ExpectedStatusCode: 400,
		},
		{
			name:               "missing user id",
			inputBody:          "title=Meeting&date=2019-09-09",
			ExpectedStatusCode: 400,
		},
	}

	for _, tc := range testCases {
		data := strings.NewReader(tc.inputBody)
		req := httptest.NewRequest("POST", "/create_event", data)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		h.CreateEvent(w, req)
		resp := w.Result()

		assert.Equal(t, tc.ExpectedStatusCode, resp.StatusCode)
	}
}
