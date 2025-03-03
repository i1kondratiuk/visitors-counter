package web

import (
	"fmt"
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

	currentVisit := &value.Visit{
		Type:  value.ResourcePath,
		Value: "homepage",
	}

	err := h.VisitLogApp.RegisterVisit(currentVisit, h.AuthApp.GetCurrentUser().Credentials.Username)

	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to log the visit; "+err.Error())
		return
	}

	uniqueVisitsNumber, err := h.VisitLogApp.GetNumberOfUsersVisitedPage(currentVisit)
	totalVisitsNumber, err := h.VisitLogApp.GetTotalVisitsNumber(currentVisit)

	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to get visit log records; "+err.Error())
		return
	}

	result := fmt.Sprintf(
		"Total number of visits: %d\nNumber of unique visitors: %d",
		totalVisitsNumber,
		uniqueVisitsNumber,
	)

	JSON(w, http.StatusOK, result)
}
