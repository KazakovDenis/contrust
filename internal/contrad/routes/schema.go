package routes

import (
	"log"
	"net/http"

	"github.com/KazakovDenis/contra/internal/contrad/scenario"
)

func Schema(w http.ResponseWriter, r *http.Request) {
	var err error

	switch r.Method {
	case http.MethodGet:
		err = getSchema(&w, r)
	case http.MethodPost:
		err = addSchema(&w, r)
	default:
		err = notAllowed(&w)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func addSchema(w *http.ResponseWriter, r *http.Request) error {
	return scenario.NewAddSchemaScenario().Execute(w, r)
}

func getSchema(w *http.ResponseWriter, r *http.Request) error {
	return scenario.NewGetSchemaScenario().Execute(w, r)
}
