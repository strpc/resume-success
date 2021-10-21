package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/strpc/resume-success/pkg/logging"

	"github.com/gorilla/mux"
)

type Handler interface {
	Register(router *mux.Router)
}

type Server struct {
	server http.Server
	logger *logging.Logger
	router *mux.Router
}

func NewServer(logger *logging.Logger, port int, router *mux.Router) *Server {
	server := &Server{
		server: http.Server{
			Addr:         fmt.Sprintf("0.0.0.0:%d", port),
			Handler:      router,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		},
		logger: logger,
		router: router,
	}
	server.initHealthCheckRoute()
	return server
}

func (s *Server) initHealthCheckRoute() {
	s.router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
