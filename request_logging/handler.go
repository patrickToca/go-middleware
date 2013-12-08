package request_logging

import (
	"log"
	"net/http"
	"time"
)

// See: https://github.com/codegangsta/martini/blob/master/logger.go

type LogRequestFunc func(*http.Request, *time.Duration)

type Handler struct {
	LogRequest LogRequestFunc
}

func (h *Handler) Process(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	started := time.Now()
	defer func() {
		elapsed := time.Since(started)
		if h.LogRequest == nil {
			h.LogRequest = DefaultLogRequest
		}

		h.LogRequest(r, &elapsed)
	}()

	next(w, r)
}

func NewHandler(logRequest LogRequestFunc) *Handler {
	return &Handler{LogRequest: logRequest}
}

func NewDefaultHandler() *Handler {
	return NewHandler(DefaultLogRequest)
}

// Helpers

func DefaultLogRequest(r *http.Request, elapsed *time.Duration) {
	log.Printf("%s: %s | Elapsed: %s\n", r.Method, r.URL.RequestURI(), elapsed.String())
}
