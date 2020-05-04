package web

import (
	"net/http"

	"github.com/i1kondratiuk/visitors-counter/application"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// UserCounterHandler ...
type UserCounterHandler struct {
	UserCounterApp application.UsersCounterApp
}

// Adds UserCounterHandler routs
func (h UserCounterHandler) AddRoutes() {
	http.HandleFunc("/homepage", h.getNumberOfVisits)
}

func (h UserCounterHandler) getNumberOfVisits(w http.ResponseWriter, r *http.Request) {
	h.UserCounterApp = application.GetUsersCounter()

	err := h.UserCounterApp.RegisterVisit(value.Visit{}, "current user username")

	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to log the visit")
		return
	}

	uniqueVisitsNumber, err := h.UserCounterApp.GetNumberOfUsersVisitedPage()

	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to get visit log records")
		return
	}

	JSON(w, http.StatusOK, uniqueVisitsNumber)
}
