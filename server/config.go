package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) apiGetConfig(w http.ResponseWriter, r *http.Request) {
	cfg, err := s.proxy.Config()
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting config: %s", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cfg); err != nil {
		http.Error(w, fmt.Sprintf("error serializing config: %s", err), http.StatusInternalServerError)
		return
	}
}

func (s *Server) apiGetConfigRaw(w http.ResponseWriter, r *http.Request) {
	cfg, err := s.proxy.Config()
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting config: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write(cfg.Body())
}
