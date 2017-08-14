package server

import (
	"fmt"
	"net/http"
)

func (s *Server) apiReload(w http.ResponseWriter, r *http.Request) {
	if err := s.proxy.Reload(); err != nil {
		http.Error(w, fmt.Sprintf("error reloading: %s", err), http.StatusInternalServerError)
		return
	}
}
