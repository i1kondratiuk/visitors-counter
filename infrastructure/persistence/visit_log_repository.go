package persistence

import (
	"database/sql"
	"errors"

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

func (r VisitLogRepositoryImpl) RegisterVisit(visit value.Visit, username string) (*entity.VisitLog, error) {
	if r.db == nil {
		return nil, errors.New("database error")
	}

	row, err := r.db.Query(
		"select id, counter from visit_log where username = ? and type = ? and value = ? limit 1",
		username,
		visit.Type,
		visit.Value,
	)

	if err != nil {
		return nil, err
	}

	var visitLogRecord *entity.VisitLog
	if row != nil {
		err = row.Scan(&visitLogRecord.Id, &visitLogRecord.Counter)
		if err != nil {
			return nil, err
		}

		_, err := r.db.Query("insert into user (name) values (?)")
		if err != nil {
			return nil, err
		}
	}

	insertedRow, err := r.db.Query("insert into visit_log (username, type, value) values (?, ?, ?)")
	if err != nil {
		return nil, err
	}

	var insertedVisitLogRecord *entity.VisitLog
	err = insertedRow.Scan(&insertedVisitLogRecord.Id, &insertedVisitLogRecord.Counter)
	if err != nil {
		return nil, err
	}

	return insertedVisitLogRecord, nil
}

func (r VisitLogRepositoryImpl) GetAllByTypeAndValue(visitType value.VisitType, visitValue string) ([]*entity.VisitLog, error) {
	if r.db == nil {
		return nil, errors.New("database error")
	}

	rows, err := r.db.Query("select value from visit_log where type = ? and value = ?", visitType, visitValue)

	if err != nil {
		return nil, err
	}

	visitLogRecords := make([]*entity.VisitLog, 0)
	for rows.Next() {
		visitLogRecord := &entity.VisitLog{}
		err = rows.Scan(
			&visitLogRecord.Id,
			&visitLogRecord.Username,
			&visitLogRecord.Counter,
			&visitLogRecord.Visit.Type,
			&visitLogRecord.Visit.Value,
		)
		if err != nil {
			return nil, err
		}
		visitLogRecords = append(visitLogRecords, visitLogRecord)
	}

	return visitLogRecords, nil
}
