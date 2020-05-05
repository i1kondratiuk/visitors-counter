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
	Signin(*value.Credentials) error
	Authorized() bool
	GetCurrentUser() *entity.User
}

// AuthImplImpl is the implementation of AuthImpl
type AuthAppImpl struct{}

// AuthImplImpl implements the AuthApp interface
var _ AuthApp = &AuthAppImpl{}

var authApp AuthApp

var authorized = false
var currentUser *entity.User

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
func (a *AuthAppImpl) Signin(credentials *value.Credentials) error {
	storedCredentials, err := repository.GetAuthRepository().GetCredentials(credentials.Username)

	if err != nil {
		return err
	}

	if err := service.GetAuthService().ComparePassword(&storedCredentials.Password, &credentials.Password); err != nil {
		return err
	}

	authorized = true

	storedUser, err := repository.GetUserRepository().GetByUsername(credentials.Username)

	if err != nil {
		return err
	}

	currentUser = storedUser

	return nil
}

func (a *AuthAppImpl) Authorized() bool {
	return authorized
}

func (a *AuthAppImpl) GetCurrentUser() *entity.User {
	return currentUser
}
