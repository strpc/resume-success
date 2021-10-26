package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/strpc/resume-success/internal/rest"

	"github.com/gorilla/mux"
	"github.com/strpc/resume-success/pkg/logging"
)

type Handler struct {
	server  *rest.Server
	logger  *logging.Logger
	service *Service
}

func NewHandler(logger *logging.Logger, service *Service, server *rest.Server) *Handler {
	return &Handler{
		server:  server,
		logger:  logger,
		service: service,
	}
}

func (h *Handler) Register(r *mux.Router) {
	h.logger.Info("Register users routes...")
	r.HandleFunc("/register", h.register()).Methods("POST")
	r.HandleFunc("/reset_password", h.resetPassword()).Methods("POST")
	r.HandleFunc("/login", h.login()).Methods("POST")
	r.HandleFunc("/logout", h.logout()).Methods("POST")
}

func (h *Handler) register() http.HandlerFunc {
	type request struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required,gte=8"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.server.ErrorResponse(w, r, http.StatusBadRequest, err)
			return
		}
		if err := h.server.ValidateStruct(req); err != nil {
			h.server.ErrorResponse(w, r, http.StatusBadRequest, err)
			return
		}
		u := &User{
			Email:    req.Email,
			Password: req.Password,
		}
		user, err := h.service.RegisterUser(u)
		h.logger.Printf("%#v", user)
		if err != nil {
			h.server.ErrorResponse(w, r, http.StatusInternalServerError, errors.New("internal error"))
			return
		}
		h.server.Response(w, r, http.StatusCreated, user)
	}
}

func (h *Handler) resetPassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		h.logger.Info("Fish for reset_password method. haha")
	}
}

func (h *Handler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		h.logger.Info("Fish for login method. haha")
	}
}

func (h *Handler) logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		h.logger.Info("Fish for logout method. haha")
	}
}
