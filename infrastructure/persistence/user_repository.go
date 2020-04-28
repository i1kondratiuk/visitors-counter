package persistence

import (
	"errors"
	"github.com/i1kondratiuk/visitors-counter/domain/entity"
)

// UserRepositoryImpl is the implementation of UserRepository
type UserRepositoryImpl struct {
	DB map[int64]*entity.User
}

// Get returns a user from the database with the id
func (r *UserRepositoryImpl) Get(id int64) (*entity.User, error) {
	if r.DB == nil {
		return nil, errors.New("database error")
	}

	if r.DB[id] == nil {
		return nil, errors.New("user not found")
	}

	return r.DB[id], nil
}

// GetAll return all users stored in database
func (r *UserRepositoryImpl) GetAll() ([]*entity.User, error) {
	if r.DB == nil {
		return nil, errors.New("database error")
	}

	users := []*entity.User{}
	for _, user := range r.DB {
		users = append(users, user)
	}

	return users, nil
}

// Save insert a user to database
func (r *UserRepositoryImpl) Save(user *entity.User) error {
	if user == nil {
		return errors.New("nil user")
	}
	if r.DB == nil {
		return errors.New("database error")
	}

	user.ID = int64(len(r.DB) + 1)
	r.DB[user.ID] = user
	return nil
}
