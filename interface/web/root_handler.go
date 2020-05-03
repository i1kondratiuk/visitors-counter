package web

import (
	"fmt"
	"log"
	"net/http"
)

// Run starts server
func Run(port int) {
	log.Printf("Server running at http://localhost:%d/", port)

	http.HandleFunc("/", root)

	new(AuthHandler).AddRoutes()
	new(UserCounterHandler).AddRoutes()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rootForm)
}

const rootForm = `
<html>
  <body>
    <h1>Visitors Counter</h1>
    <a href="/signup">Signup</a>
    </br>
    <a href="/signin">Signin</a>
  </body>
</html>
`
