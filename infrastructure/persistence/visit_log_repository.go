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

func (r VisitLogRepositoryImpl) GetVisit(visit *value.Visit, username string) (*entity.VisitLog, error) {
	if r.db == nil {
		return nil, errors.New("database error")
	}

	rows, err := r.db.Query(
		"select id, counter from visit_log where username = ? and type = ? and value = ? limit 1",
		username,
		visit.Type,
		visit.Value,
	)

	if err != nil {
		return nil, err
	}

	visitLogRecords := make([]*entity.VisitLog, 0)
	for rows.Next() {
		visitLogRecord := &entity.VisitLog{}
		err = rows.Scan(
			&visitLogRecord.Id,
			&visitLogRecord.Counter,
		)
		if err != nil {
			return nil, err
		}
		visitLogRecords = append(visitLogRecords, visitLogRecord)
	}

	if len(visitLogRecords) == 0 {
		return nil, nil
	}

	return visitLogRecords[0], nil
}

func (r VisitLogRepositoryImpl) InsertVisit(log *entity.VisitLog) (*entity.VisitLog, error) {
	_, err := r.db.Query(
		"insert into visit_log (username, type, value, counter) values (?, ?, ?, ?)",
		log.Username,
		log.Visit.Type,
		log.Visit.Value,
		log.Counter,
	)

	if err != nil {
		panic(err)
		return nil, err
	}

	return log, nil
}

func (r VisitLogRepositoryImpl) UpdateVisit(log *entity.VisitLog) (*entity.VisitLog, error) {
	_, err := r.db.Query("update visit_log set counter = ? where id = ?", log.Counter, log.Id)
	if err != nil {
		return nil, err
	}

	return log, nil
}

func (r VisitLogRepositoryImpl) GetAllByTypeAndValue(visitType *value.VisitType, visitValue string) ([]*entity.VisitLog, error) {
	if r.db == nil {
		return nil, errors.New("database error")
	}

	rows, err := r.db.Query("select id, username, counter, type, value from visit_log where type = ? and value = ?", visitType, visitValue)

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
