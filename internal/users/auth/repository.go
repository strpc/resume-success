package auth

import (
	"github.com/strpc/resume-success/pkg/clients/postgres"
	"github.com/strpc/resume-success/pkg/logging"
)

type Repository interface {
	CreateUser(u *User) (int64, error)
}

type PostgresRepository struct {
	db     *postgres.Client
	logger *logging.Logger
}

func NewPostgresRepository(logger *logging.Logger, db *postgres.Client) Repository {
	var authRepo Repository = &PostgresRepository{
		db:     db,
		logger: logger,
	}
	logger.Info("Make new db.")
	return authRepo
}

func (r *PostgresRepository) CreateUser(u *User) (int64, error) {
	var id int64
	err := r.db.QueryRowx("INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id", u.Email, u.PasswordHash).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
