package repository

import (
	"github.com/i1kondratiuk/visitors-counter/domain/entity"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// AuthRepository represents a storage of all existing ducks
type AuthRepository interface {
	GetCredentials(username string) (*value.Credentials, error)
	Insert(user *entity.User) error
}

var authRepository AuthRepository

// GetAuthRepository returns the AuthRepository
func GetAuthRepository() AuthRepository {
	return authRepository
}

// InitAuthRepository injects AuthRepository with its implementation
func InitAuthRepository(r AuthRepository) {
	authRepository = r
}
