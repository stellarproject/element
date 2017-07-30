package server

import "github.com/gorilla/mux"

func (s *Server) router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", s.genericHandler)
	r.HandleFunc("/config", s.getConfig).Methods("GET")
	r.HandleFunc("/config/raw", s.getConfigRaw).Methods("GET")
	r.HandleFunc("/frontends", s.addFrontend).Methods("POST")
	r.HandleFunc("/frontends", s.updateFrontend).Methods("PUT")
	r.HandleFunc("/frontends/{name}", s.removeFrontend).Methods("DELETE")
	r.HandleFunc("/services", s.registerService).Methods("POST")
	r.HandleFunc("/services", s.getServices).Methods("GET")
	r.HandleFunc("/reload", s.reload).Methods("POST")

	return r
}
