package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/strpc/resume-success/pkg/logging"
)

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
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

var validate *validator.Validate

func (s *Server) ValidateStruct(e interface{}) error {
	if e == nil {
		return errors.New("clean interface for validate")
	}
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) ErrorResponse(w http.ResponseWriter, r *http.Request, code int, err error) {
	if err == nil {
		err = errors.New("unknown error")
	}
	s.Response(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) Response(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			s.logger.Errorf("error response encode body %#v", data)
			s.ErrorResponse(w, r, http.StatusInternalServerError, err)
		}
	}
}
