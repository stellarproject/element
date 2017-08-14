package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ehazlett/element/proxy"
	"github.com/gorilla/mux"
)

func (s *Server) apiAddFrontend(w http.ResponseWriter, r *http.Request) {
	var frontend *proxy.Frontend
	if err := json.NewDecoder(r.Body).Decode(&frontend); err != nil {
		http.Error(w, fmt.Sprintf("invalid frontend: %s", err), http.StatusBadRequest)
		return
	}

	if err := s.proxy.AddFrontend(frontend); err != nil {
		http.Error(w, fmt.Sprintf("error adding frontend: %s", err), http.StatusInternalServerError)
		return
	}
}

func (s *Server) apiRemoveFrontend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if err := s.proxy.RemoveFrontend(name); err != nil {
		http.Error(w, fmt.Sprintf("error removing frontend: %s", err), http.StatusInternalServerError)
		return
	}
}

func (s *Server) apiUpdateFrontend(w http.ResponseWriter, r *http.Request) {
	var frontend *proxy.Frontend
	if err := json.NewDecoder(r.Body).Decode(&frontend); err != nil {
		http.Error(w, fmt.Sprintf("invalid fronend: %s", err), http.StatusBadRequest)
		return
	}

	if err := s.proxy.UpdateFrontend(frontend); err != nil {
		http.Error(w, fmt.Sprintf("error adding frontend: %s", err), http.StatusInternalServerError)
		return
	}
}
