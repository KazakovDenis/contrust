package routes

import (
	"errors"
	"net/http"

	"github.com/KazakovDenis/contrust/internal/server/local_errors"
	"github.com/KazakovDenis/contrust/internal/server/request"
	"github.com/KazakovDenis/contrust/internal/server/scenario"
)

func ProviderRouter(w http.ResponseWriter, r *http.Request) {
	httpCtx := request.NewHttpContext(&w, r)

	switch r.Method {
	case http.MethodPost:
		addProvider(httpCtx)
	default:
		httpCtx.NotAllowed()
	}
}

func addProvider(httpCtx *request.HttpContext) {
	jsonData, err := httpCtx.Json()
	if err != nil {
		httpCtx.MakeResponse(http.StatusBadRequest, "Invalid data", "text/plain")
		return
	}

	providerName, exists := jsonData["name"]
	if !exists || providerName == "" {
		httpCtx.MakeResponse(http.StatusBadRequest, "Payload must contain \"name\"", "text/plain")
		return
	}

	result, err := scenario.NewAddProviderScenario(providerName.(string)).Execute(httpCtx)
	if err == nil {
		httpCtx.MakeResponse(http.StatusOK, result, "text/plain")
		return
	}

	var writeException *local_errors.DatabaseWriteError
	switch {
	case errors.As(err, &writeException):
		httpCtx.MakeResponse(http.StatusConflict, "Already exists", "text/plain")
	default:
		httpCtx.MakeResponse(http.StatusInternalServerError, "", "text/plain")
	}
}
