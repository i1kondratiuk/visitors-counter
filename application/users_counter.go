package application

import (
	"github.com/i1kondratiuk/visitors-counter/domain/entity"
	"github.com/i1kondratiuk/visitors-counter/domain/repository"
)

// UsersCounter represents UsersCounter application to be called by interface layer
type UsersCounterApp interface {
	GetUsers() ([]*entity.User, error)
}

// UsersCounterImpl is the implementation of UsersCounter
type UsersCounterAppImpl struct{}

var usersCounterApp UsersCounterApp

// InitUsersCounter injects implementation for UsersCounter application
func InitUsersCounterApp(a UsersCounterApp) {
	usersCounterApp = a
}

// GetUsersCounter returns UsersCounter application
func GetUsersCounter() UsersCounterApp {
	return usersCounterApp
}

// UsersCounterImpl implements the UsersCounter interface
var _ UsersCounterApp = &UsersCounterAppImpl{}

// GetUsers returns users stored in repository
func (a *UsersCounterAppImpl) GetUsers() ([]*entity.User, error) {
	return repository.GetUserRepository().GetAll()
}
