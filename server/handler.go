package server

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func (s *Server) genericHandler(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"host": r.Host,
		"uri":  r.RequestURI,
	}).Debug("request")

	// TODO: check and / or configure backend container
	time.Sleep(time.Millisecond * 1000)

	// TODO: update proxy config with new backend
	time.Sleep(time.Millisecond * 1000)

	// TODO: issue redirect to host to have client re-send and connect to backend

	w.Header().Set("Location", r.RequestURI)
	w.WriteHeader(http.StatusFound)
}
