package service

import (
	"errors"
	"github.com/i1kondratiuk/visitors-counter/domain/entity"
	"github.com/i1kondratiuk/visitors-counter/domain/repository"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// UserMatchingService represents the service to find a match
type UserMatchingService interface {
	FindMatch(user *entity.User, location value.Location) (*entity.User, error)
}

// UserMatchingServiceImpl is the implementation of UserMatchingService
type UserMatchingServiceImpl struct{}

var userMatchingService UserMatchingService

// GetUserMatchingService returns a UserMatchingService
func GetUserMatchingService() UserMatchingService {
	return userMatchingService
}

// InitUserMatchingService injects UserMatchingService with its implementation
func InitUserMatchingService(m UserMatchingService) {
	userMatchingService = m
}
