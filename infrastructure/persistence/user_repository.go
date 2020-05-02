package persistence

import (
	"database/sql"
	"errors"

	"github.com/i1kondratiuk/visitors-counter/domain/entity"
	"github.com/i1kondratiuk/visitors-counter/domain/repository"
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

	row, err := r.db.Query("select id, name from users where id=?", id)
	if err != nil {
		return nil, err
	}
	user := &entity.User{}
	err = row.Scan(&user.ID, &user.Name)
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

	row, err := r.db.Query("select id, name from users where username=?", username)
	if err != nil {
		return nil, err
	}
	user := &entity.User{}
	err = row.Scan(&user.ID, &user.Name)
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
	rows, err := r.db.Query("select id, name from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*entity.User, 0)
	for rows.Next() {
		user := &entity.User{}
		err = rows.Scan(&user.ID, &user.Name)
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

	defer r.db.Close()

	_, err := r.db.Query("insert into users (name) values (?)")
	if err != nil {
		return err
	}

	return nil
}
