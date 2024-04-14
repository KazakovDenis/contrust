package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/KazakovDenis/contra/internal/contrad/scenario"
)

func addProvider(w *http.ResponseWriter, r *http.Request) {
	var jsonData map[string]string
	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		MakeResponse(w, http.StatusBadRequest, "")
		return
	}

	var providerName string
	if providerName = jsonData["name"]; providerName == "" {
		MakeResponse(w, http.StatusBadRequest, "Payload must contain \"name\"")
		return
	}

	result, err := scenario.NewAddProviderScenario(providerName).Execute(w, r)
	if err == nil {
		MakeResponse(w, http.StatusOK, result)
		return
	}

	var writeException mongo.WriteException
	switch {
	case errors.As(err, &writeException):
		MakeResponse(w, http.StatusConflict, "Already exists")
	default:
		MakeResponse(w, http.StatusInternalServerError, "")
	}
}

func Provider(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addProvider(&w, r)
	default:
		NotAllowed(&w)
	}
}
