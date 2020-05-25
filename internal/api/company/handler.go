package company

import (
	"encoding/json"
	"net/http"

	web "github.com/hrple/common/server"
)

// Handler moo
type Handler struct {
}

// RegisterHandler registers the routes for the company API
func RegisterHandler() {
	h := &Handler{}

	web.GetFunc("/companies", h.getCompanies)
}

func (h *Handler) getCompanies(w http.ResponseWriter, r *http.Request) {
	company := Company{}
	company.Name = "Company Name"

	companies := []Company{}

	companies = append(companies, company, company)

	// TODO: Abstract to separate to re-usable function
	json, err := json.Marshal(companies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
