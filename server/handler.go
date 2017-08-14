package server

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *Server) genericHandler(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"host": r.Host,
		"uri":  r.RequestURI,
	}).Debug("request")

	if err := s.connect(r.Host); err != nil {
		http.Error(w, fmt.Sprintf("error connecting to backend: %s", err), http.StatusInternalServerError)
		return
	}

	// TODO: issue redirect to host to have client re-send and connect to backend
	w.Header().Set("Location", r.RequestURI)
	w.WriteHeader(http.StatusFound)
}
