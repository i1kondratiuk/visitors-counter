package application

import (
	"github.com/i1kondratiuk/visitors-counter/domain/entity"
	"github.com/i1kondratiuk/visitors-counter/domain/repository"
	"github.com/i1kondratiuk/visitors-counter/domain/service"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// UsersCounter represents UsersCounter application to be called by interface layer
type UsersCounter interface {
	GetUser(id int64) (*entity.User, error)
	GetUsers() ([]*entity.User, error)
	AddUser(name string) error
	GetUserMatch(id int64, x, y int) (*entity.User, error)
}

// UsersCounterImpl is the implementation of UsersCounter
type UsersCounterImpl struct{}

var usersCounter UsersCounter

// InitUsersCounter injects implementation for UsersCounter application
func InitUsersCounter(t UsersCounter) {
	usersCounter = t
}
// GetUsersCounter returns UsersCounter application
func GetUsersCounter() UsersCounter {
	return usersCounter
}

// UsersCounterImpl implements the UsersCounter interface
var _ UsersCounter = &UsersCounterImpl{}

// GetUser returns user with the given id
func (t *UsersCounterImpl) GetUser(id int64) (*entity.User, error) {
	return repository.GetUserRepository().Get(id)
}

// GetUsers returns users stored in repository
func (t *UsersCounterImpl) GetUsers() ([]*entity.User, error) {
	return repository.GetUserRepository().GetAll()
}

// AddUser inserts new user with the given name to repository
func (t *UsersCounterImpl) AddUser(name string) error {
	return repository.GetUserRepository().Save(&entity.User{
		Name: name,
	})
}
