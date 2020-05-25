package status

import (
	"encoding/json"
	"net/http"

	web "github.com/hrple/common/server"
)

// Handler
type Handler struct {
}

// RegisterHandler registers the routes for the status API
func RegisterHandler() {
	h := &Handler{}

	web.GetFunc("/status", h.status)
}

func (h *Handler) status(w http.ResponseWriter, r *http.Request) {
	status := Status{}
	status.IsRunning = true

	json, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
