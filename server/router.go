package server

import "github.com/gorilla/mux"

func (s *Server) router() *mux.Router {
	r := mux.NewRouter()
	// generic
	r.HandleFunc("/", s.genericHandler)
	// proxy
	r.HandleFunc("/config", s.apiGetConfig).Methods("GET")
	r.HandleFunc("/config/raw", s.apiGetConfigRaw).Methods("GET")
	r.HandleFunc("/frontends", s.apiAddFrontend).Methods("POST")
	r.HandleFunc("/frontends", s.apiUpdateFrontend).Methods("PUT")
	r.HandleFunc("/frontends/{name}", s.apiRemoveFrontend).Methods("DELETE")
	r.HandleFunc("/reload", s.apiReload).Methods("POST")

	return r
}
