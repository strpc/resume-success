package users

import "github.com/strpc/resume-success/pkg/logging"

type UserService struct {
	logger     *logging.Logger
	repository *UserRepository
}

func NewUserService(logger *logging.Logger, repository *UserRepository) *UserService {
	return &UserService{
		logger:     logger,
		repository: repository,
	}
}
