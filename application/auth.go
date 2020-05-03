package application

import (
	"github.com/i1kondratiuk/visitors-counter/domain/entity"
	"github.com/i1kondratiuk/visitors-counter/domain/repository"
	"github.com/i1kondratiuk/visitors-counter/domain/service"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// AuthApp...
type AuthApp interface {
	Signup(*entity.User) error
	Signin(*value.Credentials) (*entity.User, error)
}

// AuthImplImpl is the implementation of AuthImpl
type AuthAppImpl struct{}

// AuthImplImpl implements the AuthApp interface
var _ AuthApp = &AuthAppImpl{}

var authApp AuthApp

// InitAuthApp injects implementation for AuthApp application
func InitAuthApp(a AuthApp) {
	authApp = a
}

// GetAuthApp returns AuthApp application
func GetAuthApp() AuthApp {
	return authApp
}

// Signup inserts a user
func (a *AuthAppImpl) Signup(user *entity.User) error {
	err := repository.GetAuthRepository().Insert(user)

	if err != nil {
		panic(err.Error())
	}

	return nil
}

// Signin checks for credentials to match
func (a *AuthAppImpl) Signin(c *value.Credentials) (*entity.User, error) {
	u, err := service.GetAuthService().FindMatch(c)

	if err == nil {
		return nil, err
	}

	return u, nil
}
