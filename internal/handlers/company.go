package handlers

import (
	"fmt"
	"net/http"

	web "github.com/hrple/common/server"
)

// CompanyHandler
type CompanyHandler struct {
}

// RegisterCompanyHandler registers the routes for the company API
func RegisterCompanyHandler() {
	h := &CompanyHandler{}

	web.GetFunc("/hello", h.hello)
}

func (h *CompanyHandler) hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hello %v \n", r.URL)
}
