package routes

import (
	"net/http"

	"github.com/KazakovDenis/contra/internal/common/log"
	"github.com/KazakovDenis/contra/internal/contrad/scenario"
)

func Schema(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getSchema(&w, r)
	case http.MethodPost:
		addSchema(&w, r)
	default:
		NotAllowed(&w)
	}
}

func addSchema(w *http.ResponseWriter, r *http.Request) {
	err := scenario.NewAddSchemaScenario().Execute(w, r)
	if err != nil {
		log.Error("%s", err)
	}
}

func getSchema(w *http.ResponseWriter, r *http.Request) {
	err := scenario.NewGetSchemaScenario().Execute(w, r)
	if err != nil {
		log.Error("%s", err)
	}
}
