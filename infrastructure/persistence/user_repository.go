package persistence

import (
	"database/sql"
	"errors"

	"github.com/i1kondratiuk/visitors-counter/domain/entity"
	"github.com/i1kondratiuk/visitors-counter/domain/repository"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// UserRepositoryImpl is the implementation of UserRepository
type UserRepositoryImpl struct {
	db *sql.DB
}

// UserRepositoryImpl implements the domain UserRepository interface
var _ repository.UserRepository = &UserRepositoryImpl{}

// NewUserRepository returns initialized UserRepositoryImpl
func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

// Get returns a user from the database with the id
func (r *UserRepositoryImpl) GetById(id int64) (*entity.User, error) {
	if r.db == nil {
		return nil, errors.New("database error")
	}

	row, err := r.db.Query("select id, name from user where id=?", id)
	if err != nil {
		return nil, err
	}
	user := &entity.User{}
	err = row.Scan(&user.Id, &user.Name)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Get returns a user from the database with the username
func (r *UserRepositoryImpl) GetByUsername(username string) (*entity.User, error) {
	if r.db == nil {
		return nil, errors.New("database error")
	}

	rows, err := r.db.Query("select id, name from user where username=?", username)
	if err != nil {
		return nil, err
	}

	user := &entity.User{Credentials: value.Credentials{Username: username}}
	rows.Next()
	err = rows.Scan(&user.Id, &user.Name)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetAll returns list of domain.User
func (r *UserRepositoryImpl) GetAll() ([]*entity.User, error) {
	if r.db == nil {
		return nil, errors.New("database error")
	}

	rows, err := r.db.Query("SELECT id, name FROM user")

	if err != nil {
		return nil, err
	}

	users := make([]*entity.User, 0)
	for rows.Next() {
		user := &entity.User{}
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Save saves domain.User to storage
func (r *UserRepositoryImpl) Save(user *entity.User) error {
	if user == nil {
		return errors.New("nil user")
	}
	if r.db == nil {
		return errors.New("database error")
	}

	_, err := r.db.Query("insert into user (name) values (?)")
	if err != nil {
		return err
	}

	return nil
}
