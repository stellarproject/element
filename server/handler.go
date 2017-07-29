package server

import (
	"net/http"

	"github.com/ehazlett/element/version"
)

func (s *Server) getRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Content-Server", "element "+version.FullVersion())
	w.WriteHeader(http.StatusOK)
}
