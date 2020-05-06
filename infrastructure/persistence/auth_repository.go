package persistence

import (
	"database/sql"
	"errors"

	"github.com/i1kondratiuk/visitors-counter/domain/entity"
	"github.com/i1kondratiuk/visitors-counter/domain/repository"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// UserRepositoryImpl is the implementation of UserRepository
type AuthRepositoryImpl struct {
	db *sql.DB
}

// UserRepositoryImpl implements the domain AuthRepository interface
var _ repository.AuthRepository = &AuthRepositoryImpl{}

// NewUserRepository returns initialized UserRepositoryImpl
func NewAuthRepository(db *sql.DB) repository.AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

// Save saves domain.User to storage
func (r *AuthRepositoryImpl) GetCredentials(username string) (*value.Credentials, error) {
	storedCredentials := &value.Credentials{}
	rows, err := r.db.Query("select username, password from user where username = ? limit 1", username)
	if rows == nil {
		return nil, errors.New("no records found")
	}
	if err != nil {
		return nil, errors.New("no records found")
	}

	rows.Next()
	err = rows.Scan(
		&storedCredentials.Username,
		&storedCredentials.Password,
	)

	if err != nil {
		return nil, errors.New("no user with such username found")
	}

	return storedCredentials, nil
}

// Save saves domain.User to storage
func (r *AuthRepositoryImpl) Insert(user *entity.User) error {
	if user == nil {
		return errors.New("nil user")
	}

	if r.db == nil {
		return errors.New("database error")
	}

	_, err := r.db.Query(
		"insert into user (name, username, password) values (?, ?, ?)",
		user.Name,
		user.Credentials.Username,
		user.Credentials.Password)

	if err != nil {
		return err
	}

	return nil
}
