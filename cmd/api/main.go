package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/strpc/resume-success/internal/users"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
	cfg "github.com/strpc/resume-success/internal/config"
	"github.com/strpc/resume-success/internal/rest"
	"github.com/strpc/resume-success/pkg/logging"
)

func main() {
	config := cfg.GetConfig()
	logger := logging.NewLogger(config.App.LogLevel, config.App.LogType)

	router := mux.NewRouter()

	userRepo := users.NewUserRepository(logger, "postgres maafaka")
	userService := users.NewUserService(logger, userRepo)

	usersHandler := users.NewUserHandler(logger, userService)
	usersHandler.Register(router)

	server := rest.NewServer(config.App.Port, router)

	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logger.Infof("Server Started on %d port", config.App.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("Server Shutting Down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("error occured on server shutting down: %s", err)
	}
}
