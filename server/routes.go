package server

import (
	"github.com/gorilla/mux"
)

func (s *Server) router() (*mux.Router, error) {
	r := mux.NewRouter()
	r.HandleFunc("/", s.getRequestHandler).Methods("GET")

	return r, nil
}
