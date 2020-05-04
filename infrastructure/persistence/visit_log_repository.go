package persistence

import (
	"database/sql"

	"github.com/i1kondratiuk/visitors-counter/domain/entity"
	"github.com/i1kondratiuk/visitors-counter/domain/repository"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// VisitLogRepositoryImpl is the implementation of VisitLogRepository
type VisitLogRepositoryImpl struct {
	db *sql.DB
}

// VisitLogRepositoryImpl implements the domain VisitLogRepository interface
var _ repository.VisitLogRepository = &VisitLogRepositoryImpl{}

// VisitLogRepository returns initialized VisitLogRepositoryImpl
func NewVisitLogRepository(db *sql.DB) repository.VisitLogRepository {
	return &VisitLogRepositoryImpl{db: db}
}

func (v VisitLogRepositoryImpl) RegisterVisit(visit value.Visit, username string) error {
	panic("implement me")
}

func (v VisitLogRepositoryImpl) GetByTypeAndUsername(visitType value.VisitType) ([]*entity.VisitLog, error) {
	panic("implement me")
}
