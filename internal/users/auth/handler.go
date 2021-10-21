package auth

import (
	"encoding/json"
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
	type RequestBody struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required,gte=8"`
	}
	var b RequestBody

	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			h.server.ErrorResponse(w, r, http.StatusBadRequest, err)
			return
		}
		if err := h.server.ValidateStruct(b); err != nil {
			h.server.ErrorResponse(w, r, http.StatusBadRequest, err)
			return
		}

		user, err := h.service.RegisterUser(b.Email, b.Password)
		if err != nil {
			h.server.ErrorResponse(w, r, http.StatusInternalServerError, err)
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
