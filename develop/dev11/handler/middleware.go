package handler

import (
	"log"
	"net/http"
	"time"
)

func loggingRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(started))
	}
}
