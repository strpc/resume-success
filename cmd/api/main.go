package main

import (
	"context"
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

	psqlClient, err := postgres.NewClient(
		config.DB.Postgres.Host,
		config.DB.Postgres.User,
		config.DB.Postgres.Password,
		config.DB.Postgres.DBName,
		config.DB.Postgres.SSLMode,
		config.DB.Postgres.Port,
	)
	if err != nil {
		logger.Fatalf("PostgreSQL connection error. %s", err.Error())
	}
	defer func() {
		logger.Info("Closing PostgreSQL connection...")
		if err := psqlClient.Close(); err != nil {
			logger.Fatalf("Error closing PostgreSQL connection. Error: %s", err.Error())
		}
		logger.Info("PostgreSQL connection closed.")
	}()

	var server *rest.Server

	authHandler := getAuthHandler(psqlClient, logger, server)
	authRouter := router.PathPrefix("/auth").Subrouter()
	authHandler.Register(authRouter)

	server = rest.NewServer(logger, config.App.Port, router)

	go func() {
		if err := server.Start(); err != nil {
			logger.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logger.Infof("Server Started on %d port", config.App.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("Server Shutting Down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("error occured on server shutting down: %s", err)
	}
}

func getAuthHandler(pgClient *postgres.Client, logger *logging.Logger, server *rest.Server) *auth.Handler {
	authRepo := auth.NewPostgresRepository(logger, pgClient)
	authService := auth.NewService(logger, authRepo)
	authHandler := auth.NewHandler(logger, authService, server)
	return authHandler
}
