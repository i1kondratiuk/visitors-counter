package application

import (
	"github.com/i1kondratiuk/visitors-counter/domain/repository"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// UsersCounter represents UsersCounter application to be called by interface layer
type UsersCounterApp interface {
	RegisterVisit(visit value.Visit, username string) error
	GetNumberOfUsersVisitedPage() (int, error)
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

func (a *UsersCounterAppImpl) RegisterVisit(visit value.Visit, username string) error {
	err := repository.GetVisitLogRepository().RegisterVisit(visit, username)

	if err != nil {
		return err
	}

	return nil
}

func (a *UsersCounterAppImpl) GetNumberOfUsersVisitedPage() (int, error) {
	logs, err := repository.GetVisitLogRepository().GetByTypeAndUsername(value.ResourcePath)

	if err != nil {
		panic(err.Error())
	}

	return len(logs), err
}
