package repository

import (
	"github.com/i1kondratiuk/visitors-counter/domain/entity"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// VisitLogRepository represents a storage of all visit logs
type VisitLogRepository interface {
	GetVisit(visit *value.Visit, username string) (*entity.VisitLog, error)
	InsertVisit(*entity.VisitLog) (*entity.VisitLog, error)
	UpdateVisit(*entity.VisitLog) (*entity.VisitLog, error)
	GetAllByTypeAndValue(visitType *value.VisitType, visitValue string) ([]*entity.VisitLog, error)
}

var visitLogRepository VisitLogRepository

// GetVisitLogRepository returns the VisitLogRepository
func GetVisitLogRepository() VisitLogRepository {
	return visitLogRepository
}

// InitUserRepository injects UserRepository with its implementation
func InitVisitLogRepository(r VisitLogRepository) {
	visitLogRepository = r
}
