package auth

import "github.com/strpc/resume-success/pkg/logging"

type Service struct {
	logger     *logging.Logger
	repository *Repository
}

func NewService(logger *logging.Logger, repository *Repository) *Service {
	return &Service{
		logger:     logger,
		repository: repository,
	}
}
