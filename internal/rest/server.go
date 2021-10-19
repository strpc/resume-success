package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Handler interface {
	Register(router *mux.Router)
}

type Server struct {
	server http.Server
}

func NewServer(port int, router *mux.Router) *Server {
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(map[string]bool{"pong": true})
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		_, _ = w.Write(jsonResp)
	})
	return &Server{
		server: http.Server{
			Addr:         fmt.Sprintf("0.0.0.0:%d", port),
			Handler:      router,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
