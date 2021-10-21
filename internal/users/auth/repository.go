package auth

import (
	"github.com/strpc/resume-success/pkg/clients/postgres"
	"github.com/strpc/resume-success/pkg/logging"
)

type Repository interface {
	Create() error
}

type PostgresRepository struct {
	db     *postgres.Client
	logger *logging.Logger
}

func NewPostgresRepository(logger *logging.Logger, db *postgres.Client) *Repository {
	var authRepo Repository = &PostgresRepository{
		db:     db,
		logger: logger,
	}
	logger.Info("Make new db.")
	return &authRepo
}

func (r *PostgresRepository) Create() error {
	r.logger.Info("Method of create yoba")
	return nil
}
