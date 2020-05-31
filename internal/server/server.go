// Package server contains logic for bringing up the server and creating routes.
package server

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const serverAddress = "0.0.0.0:80"

// Server is a wrapper for the http.Server type.
type Server struct {
	*http.Server
}

// Serve creates a new Server and starts serving.
func Serve() {
	s := NewServer()
	s.Start()
}

// Start starts the execution of a Server.
func (s *Server) Start() {
	log.Infoln("Starting server...")
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

// NewServer creates a new server with the BaseRouter.
func NewServer() Server {
	s := Server{
		&http.Server{
			Handler:      BaseRouter,
			Addr:         serverAddress,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}

	return s
}