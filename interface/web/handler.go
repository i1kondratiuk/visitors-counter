package web

import (
	"fmt"
	"log"
	"net/http"
)

// Run starts server
func Run(port int) {
	log.Printf("Server running at http://localhost:%d/", port)

	new(AuthHandler).AddRoutes()
	new(UserCounterHandler).AddRoutes()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
