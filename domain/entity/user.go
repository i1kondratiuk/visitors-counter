package entity

import "github.com/i1kondratiuk/visitors-counter/domain/value"

// User represents a user entity stored in repository
type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Credentials value.Credentials
}
