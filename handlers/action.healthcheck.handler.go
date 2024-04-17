package handlers

import (
	"demo-shop-manager/responses"
	"net/http"
)

func handle_healthcheck(w http.ResponseWriter, r *http.Request) {
	// Prepare the response
	response := responses.NewHealthCheckResponse("ok")

	// Send the response
	responses.SendResponse(r, w, http.StatusOK, response)
}

func init() {
	GetRegistry().AddHandler("/_action/healthcheck", []string{"GET", "POST"}, handle_healthcheck)
}
