package web

import (
	"encoding/json"
	"net/http"

	"github.com/i1kondratiuk/visitors-counter/application"
	"github.com/i1kondratiuk/visitors-counter/domain/entity"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// Handler user handler
type AuthHandler struct {
	AuthApp application.AuthApp
}

// AddRoutes adds AuthHandler routs
func (h AuthHandler) AddRoutes() {
	http.HandleFunc("/signup", h.Signup)
	http.HandleFunc("/signin", h.Signin)
}

func (h AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var u entity.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		Error(w, http.StatusBadRequest, err, "failed to parse request")
		return
	}

	if err := h.AuthApp.Signup(&u); err != nil {
		Error(w, http.StatusInternalServerError, err, "failed to create user")
		return
	}

	JSON(w, http.StatusCreated, nil)
}

func (h AuthHandler) Signin(w http.ResponseWriter, r *http.Request) {

	var c value.Credentials
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		Error(w, http.StatusBadRequest, err, "failed to parse request")
		return
	}

	u, err := h.AuthApp.Signin(&c)

	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to signin")
		return
	}

	JSON(w, http.StatusOK, u)
}
