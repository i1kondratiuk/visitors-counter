package web

import (
	"net/http"

	"github.com/i1kondratiuk/visitors-counter/application"
	"github.com/i1kondratiuk/visitors-counter/domain/value"
)

// VisitLogAppHandler ...
type VisitLogAppHandler struct {
	AuthApp     application.AuthApp
	VisitLogApp application.VisitLogApp
}

// Adds UserCounterHandler routs
func (h VisitLogAppHandler) AddRoutes() {
	http.HandleFunc("/homepage", h.getNumberOfVisits)
}

func (h VisitLogAppHandler) getNumberOfVisits(w http.ResponseWriter, r *http.Request) {
	h.AuthApp = application.GetAuthApp()
	h.VisitLogApp = application.GetVisitLogApp()

	currentVisit := value.Visit{
		Type:  value.ResourcePath,
		Value: "homepage",
	}

	err := h.VisitLogApp.RegisterVisit(currentVisit, h.AuthApp.GetCurrentUser().Credentials.Username)

	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to log the visit")
		return
	}

	uniqueVisitsNumber, err := h.VisitLogApp.GetNumberOfUsersVisitedPage(currentVisit)

	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to get visit log records")
		return
	}

	JSON(w, http.StatusOK, uniqueVisitsNumber)
}
