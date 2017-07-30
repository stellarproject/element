package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ehazlett/element/proxy"
)

func (s *Server) updateFrontend(w http.ResponseWriter, r *http.Request) {
	var frontend *proxy.Frontend
	if err := json.NewDecoder(r.Body).Decode(&frontend); err != nil {
		http.Error(w, fmt.Sprintf("invalid fronend: %s", err), http.StatusBadRequest)
		return
	}

	if err := s.proxy.UpdateFrontend(frontend); err != nil {
		http.Error(w, fmt.Sprintf("error adding frontend: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
