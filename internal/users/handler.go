package users

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/strpc/resume-success/pkg/logging"
)

type UserHandler struct {
	logger  *logging.Logger
	service *UserService
}

func NewUserHandler(logger *logging.Logger, service *UserService) *UserHandler {
	return &UserHandler{
		logger:  logger,
		service: service,
	}
}

func (h *UserHandler) Register(r *mux.Router) {
	h.logger.Info("Register users routes...")
	s := r.PathPrefix("/users").Subrouter()
	s.HandleFunc("/register", h.register()).Methods("POST")
	s.HandleFunc("/reset_password", h.resetPassword()).Methods("POST")
	s.HandleFunc("/login", h.login()).Methods("POST")
	s.HandleFunc("/logout", h.logout()).Methods("POST")
}

func (h *UserHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		h.logger.Info("Fish for register method. haha")
	}
}

func (h *UserHandler) resetPassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		h.logger.Info("Fish for reset_password method. haha")
	}
}

func (h *UserHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		h.logger.Info("Fish for login method. haha")
	}
}

func (h *UserHandler) logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		h.logger.Info("Fish for logout method. haha")
	}
}
