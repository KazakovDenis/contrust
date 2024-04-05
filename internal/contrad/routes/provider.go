package routes

import (
	"log"
	"net/http"

	"github.com/KazakovDenis/contra/internal/contrad/scenario"
)

func addProvider(w *http.ResponseWriter, r *http.Request) error {
	return scenario.NewAddProviderScenario().Execute(w, r)
}

func Provider(w http.ResponseWriter, r *http.Request) {
	var err error

	switch r.Method {
	case http.MethodPost:
		err = addProvider(&w, r)
	default:
		err = notAllowed(&w)
	}

	if err != nil {
		log.Fatal(err)
	}
}
