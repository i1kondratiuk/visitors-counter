package web

import (
	"net/http"

	"github.com/i1kondratiuk/visitors-counter/application"
)

// UserCounterHandler ...
type UserCounterHandler struct {
	UserCounterApp application.UsersCounterApp
}

// Adds UserCounterHandler routs
func (h UserCounterHandler) AddRoutes() {
	http.HandleFunc("/getUsers", h.getUsers)
}

func (h UserCounterHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	us, err := h.UserCounterApp.GetUsers()

	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to get users")
		return
	}

	JSON(w, http.StatusOK, us)
}
