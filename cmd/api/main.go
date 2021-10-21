package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/strpc/resume-success/internal/users/auth"

	"github.com/gorilla/mux"
	cfg "github.com/strpc/resume-success/internal/config"
	"github.com/strpc/resume-success/internal/rest"
	"github.com/strpc/resume-success/internal/rest/middleware"
	"github.com/strpc/resume-success/pkg/clients/postgres"
	"github.com/strpc/resume-success/pkg/logging"
)

func main() {
	config := cfg.GetConfig()
	logger := logging.NewLogger(config.App.LogLevel, config.App.LogType)

	router := mux.NewRouter().PathPrefix("/api").Subrouter()
	loggingMiddleware := middleware.NewLoggingMiddleware(logger)
	router.Use(loggingMiddleware.Middleware)

	psqlConfig := config.DB.Postgres
	psqlClient, err := postgres.NewClient(
		logger,
		psqlConfig.Host,
		psqlConfig.User,
		psqlConfig.Password,
		psqlConfig.DBName,
		psqlConfig.SSLMode,
		psqlConfig.Port,
	)
	if err != nil {
		logger.Fatalf("Postgresql connection error. %s", err.Error())
	}
	userRepo := auth.NewPostgresRepository(logger, psqlClient)
	userService := auth.NewService(logger, userRepo)

	authHandler := auth.NewHandler(logger, userService)
	authRouter := router.PathPrefix("/auth").Subrouter()
	authHandler.Register(authRouter)

	server := rest.NewServer(logger, config.App.Port, router)

	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("error occured while running http server: %s", err.Error())
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
