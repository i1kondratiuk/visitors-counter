package repository

import (
	"github.com/i1kondratiuk/visitors-counter/domain/entity"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// VisitLogRepository represents a storage of all visit logs
type VisitLogRepository interface {
	RegisterVisit(visit value.Visit, username string) error
	GetByTypeAndUsername(value.VisitType) ([]*entity.VisitLog, error)
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
