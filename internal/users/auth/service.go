package auth

import (
	"fmt"

	"github.com/strpc/resume-success/pkg/logging"
	"golang.org/x/crypto/bcrypt"
)

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

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (s *Service) RegisterUser(email, password string) (User, error) {
	e, _ := encryptString(password)
	fmt.Println(e)
	return User{}, nil
}
