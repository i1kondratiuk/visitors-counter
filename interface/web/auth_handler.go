package web

import (
	"errors"
	"fmt"
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
	switch r.Method {
	case "GET":
		fmt.Fprint(w, signupForm)
	case "POST":

		u := entity.User{
			Name: r.FormValue("name"),
			Credentials: value.Credentials{
				Username: r.FormValue("username"),
				Password: r.FormValue("password"),
			},
		}

		if u.Name == "" || u.Credentials.Username == "" || u.Credentials.Password == "" {
			Error(w, http.StatusBadRequest, errors.New("empty input"), "all fields should be populated")
			return
		}

		h.AuthApp = application.GetAuthApp()
		if err := h.AuthApp.Signup(&u); err != nil {
			Error(w, http.StatusInternalServerError, err, "failed to create the user")
			return
		}

		http.Redirect(w, r, "/homepage", http.StatusMovedPermanently)
	default:
		fmt.Fprintf(w, "only GET and POST methods are supported")
	}
}

const signupForm = `
<html>
  <body>
    <h1>Signup</h1>
    <form action="/signup" method="post">
      <p>
        <label for="name">Name:</label>
        <input id="name" type="text" name="name">
      </p>
      <p>
        <label for="username">Username:</label>
        <input id="username" type="text" name="username">
      </p>
      <p>
        <label for="password">Password:</label>
        <input in="password" type="text" name="password">
      </p>
        <input value="Submit" type="submit">
    </form>
  </body>
</html>
`

func (h AuthHandler) Signin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, signinForm)
	case "POST":
		credentials := &value.Credentials{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		if credentials.Username == "" || credentials.Password == "" {
			Error(w, http.StatusBadRequest, errors.New("empty input"), "all fields should be populated")
			return
		}

		err := h.AuthApp.Signin(credentials)

		if err != nil {
			Error(w, http.StatusNotFound, err, "failed to signin")
			return
		}

		http.Redirect(w, r, "/homepage", http.StatusMovedPermanently)
	default:
		fmt.Fprintf(w, "only GET and POST methods are supported")
	}
}

const signinForm = `
<html>
  <body>
    <h1>Signin</h1>
    <form action="/signin" method="post">
      <p>
        <label for="name">Name:</label>
        <input id="name" type="text" name="name">
      </p>
      <p>
        <label for="username">Username:</label>
        <input id="username" type="text" name="username">
      </p>
      <p>
        <label for="password">Password:</label>
        <input in="password" type="text" name="password">
      </p>
        <input value="Submit" type="submit">
    </form>
  </body>
</html>
`
