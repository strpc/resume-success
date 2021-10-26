package auth

import (
	"github.com/strpc/resume-success/pkg/logging"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	logger     *logging.Logger
	repository Repository
}

func NewService(logger *logging.Logger, repository Repository) *Service {
	return &Service{
		logger:     logger,
		repository: repository,
	}
}

func encryptPassword(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func decryptPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (s *Service) RegisterUser(u *User) (*User, error) {
	h, err := encryptPassword(u.Password)
	if err != nil {
		s.logger.Error(err)
		return &User{}, err
	}
	u.PasswordHash = h
	id, err := s.repository.CreateUser(u)
	if err != nil {
		s.logger.Error(err)
		return &User{}, err
	}
	u.ID = id
	return u, nil
}
