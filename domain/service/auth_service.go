package service

import (
	"errors"

	"github.com/i1kondratiuk/visitors-counter/domain/entity"
	"github.com/i1kondratiuk/visitors-counter/domain/repository"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// AuthService represents the service to find a match
type AuthService interface {
	FindMatch(*value.Credentials) (*entity.User, error)
}

// AuthServiceImpl is the implementation of UserMatchingService
type AuthServiceImpl struct{}

var authService AuthService

// GetAuthService returns a AuthService
func GetAuthService() AuthService {
	return authService
}

// InitAuthService injects AuthServiceService with its implementation
func InitAuthService(a AuthService) {
	authService = a
}

// FindMatch finds a match for a user based on its Credentials
func (a *AuthServiceImpl) FindMatch(c *value.Credentials) (*entity.User, error) {
	retrieved, err := repository.GetAuthRepository().GetCredentials(c.Username)

	if err != nil {
		return nil, errors.New("user credentials not found")
	}

	if retrieved.Password != c.Password {
		return nil, errors.New("incorrect password")
	}

	u, err := repository.GetUserRepository().GetByUsername(c.Username)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return u, nil
}
