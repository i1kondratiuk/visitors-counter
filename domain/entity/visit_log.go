package entity

import "github.com/i1kondratiuk/visitors-counter/domain/value"

// VisitLog represents a VisitLog entity stored in repository
type VisitLog struct {
	ID       int64       `json:"id"`
	Username string      `json:"username"`
	Counter  int         `json:"counter"`
	Visit    value.Visit `json:"visit"`
}
