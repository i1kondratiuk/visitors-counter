package entity

// User represents a user entity stored in repository
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
