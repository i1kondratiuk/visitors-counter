package repository

import (
	"github.com/i1kondratiuk/visitors-counter/domain/entity"
)

// UserRepository represents a storage of all existing users
type UserRepository interface {
	Get(ID int64) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Save(user *entity.User) error
}

var userRepository UserRepository

// GetUserRepository returns the UserRepository
func GetUserRepository() UserRepository {
	return userRepository
}

// InitUserRepository injects UserRepository with its implementation
func InitUserRepository(r UserRepository) {
	userRepository = r
}
