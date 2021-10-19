package users

import "github.com/strpc/resume-success/pkg/logging"

type UserRepository interface {
	Create() error
}

type UserPostgresRepository struct {
	db     string
	logger *logging.Logger
}

func NewUserRepository(logger *logging.Logger, db string) *UserRepository {
	var userRepo UserRepository = &UserPostgresRepository{
		db:     db,
		logger: logger,
	}
	logger.Info("Make new db.")
	return &userRepo
}

func (r *UserPostgresRepository) Create() error {
	r.logger.Info("Method of create yoba")
	return nil
}
