package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/strpc/resume-success/pkg/logging"
)

type Handler struct {
	logger  *logging.Logger
	service *Service
}

func NewHandler(logger *logging.Logger, service *Service) *Handler {
	return &Handler{
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
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		h.logger.Info("Fish for register method. haha")
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
