package routes

import (
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	http2 "github.com/KazakovDenis/contra/internal/contrad/http"
	"github.com/KazakovDenis/contra/internal/contrad/scenario"
)

func addProvider(httpCtx *http2.HttpContext) {
	jsonData, err := httpCtx.Json()
	if err != nil {
		httpCtx.MakeResponse(http.StatusBadRequest, "Invalid data")
		return
	}

	var providerName string
	if providerName = jsonData["name"].(string); providerName == "" {
		httpCtx.MakeResponse(http.StatusBadRequest, "Payload must contain \"name\"")
		return
	}

	result, err := scenario.NewAddProviderScenario(providerName).Execute(httpCtx)
	if err == nil {
		httpCtx.MakeResponse(http.StatusOK, result)
		return
	}

	var writeException mongo.WriteException
	switch {
	case errors.As(err, &writeException):
		httpCtx.MakeResponse(http.StatusConflict, "Already exists")
	default:
		httpCtx.MakeResponse(http.StatusInternalServerError, "")
	}
}

func Provider(w http.ResponseWriter, r *http.Request) {
	httpCtx := http2.NewHttpContext(&w, r)

	switch r.Method {
	case http.MethodPost:
		addProvider(httpCtx)
	default:
		http2.NotAllowed(&w)
	}
}
