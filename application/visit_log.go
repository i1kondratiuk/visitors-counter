package application

import (
	"github.com/i1kondratiuk/visitors-counter/domain/repository"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// VisitLogApp represents VisitLogApp application to be called by interface layer
type VisitLogApp interface {
	RegisterVisit(visit value.Visit, username string) error
	GetNumberOfUsersVisitedPage(visit value.Visit) (int, error)
}

// VisitLogAppImpl is the implementation of UsersCounter
type VisitLogAppImpl struct{}

var visitLogAppApp VisitLogApp

// InitVisitLogApp injects implementation for VisitLogApp application
func InitVisitLogApp(a VisitLogApp) {
	visitLogAppApp = a
}

// GetVisitLogApp returns VisitLogApp application
func GetVisitLogApp() VisitLogApp {
	return visitLogAppApp
}

// VisitLogAppImpl implements the VisitLogApp interface
var _ VisitLogApp = &VisitLogAppImpl{}

func (a *VisitLogAppImpl) RegisterVisit(visit value.Visit, username string) error {
	_, err := repository.GetVisitLogRepository().RegisterVisit(visit, username)

	if err != nil {
		return err
	}

	return nil
}

func (a *VisitLogAppImpl) GetNumberOfUsersVisitedPage(visit value.Visit) (int, error) {
	logs, err := repository.GetVisitLogRepository().GetAllByTypeAndValue(visit.Type, visit.Value)

	if err != nil {
		panic(err.Error())
	}

	return len(logs), err
}
