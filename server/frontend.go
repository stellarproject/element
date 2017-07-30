package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ehazlett/element/proxy"
	"github.com/gorilla/mux"
)

func (s *Server) addFrontend(w http.ResponseWriter, r *http.Request) {
	var frontend *proxy.Frontend
	if err := json.NewDecoder(r.Body).Decode(&frontend); err != nil {
		http.Error(w, fmt.Sprintf("invalid fronend: %s", err), http.StatusBadRequest)
		return
	}

	if err := s.proxy.AddFrontend(frontend); err != nil {
		http.Error(w, fmt.Sprintf("error adding frontend: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) removeFrontend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if err := s.proxy.RemoveFrontend(name); err != nil {
		http.Error(w, fmt.Sprintf("error removing frontend: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

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
