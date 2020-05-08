package service

import (
	"errors"
)

// AuthService represents the service to find a match
type AuthService interface {
	ComparePassword(storedPassword *string, password *string) error
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
func (a *AuthServiceImpl) ComparePassword(storedPassword *string, password *string)  error {
	if *storedPassword != *password {
		return errors.New("incorrect password")
	}

	return nil
}
