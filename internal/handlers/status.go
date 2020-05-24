package handlers

import (
	"fmt"
	"net/http"

	web "github.com/hrple/common/server"
)

// StatusHandler
type StatusHandler struct {
}

// RegisterStatusHandler registers the routes for the status API
func RegisterStatusHandler() {
	h := &StatusHandler{}

	web.GetFunc("/status", h.status)
}

func (h *StatusHandler) status(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Running TRUE (Temporay Message)")
}
