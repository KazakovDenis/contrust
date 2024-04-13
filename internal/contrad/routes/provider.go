package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KazakovDenis/contra/internal/contrad/scenario"
)

func addProvider(w *http.ResponseWriter, r *http.Request) error {
	var jsonData map[string]string
	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		return err
	}
	return scenario.NewAddProviderScenario(jsonData["name"]).Execute(w, r)
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
