package routes

import (
	"log"
	"net/http"
)

func Schema(w http.ResponseWriter, r *http.Request) {
	var err error

	switch r.Method {
	case http.MethodGet:
		err = getSchema(&w)
	case http.MethodPost:
		err = addSchema(&w)
	default:
		err = notAllowed(&w)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func addSchema(w *http.ResponseWriter) error {
	return nil
}

func getSchema(w *http.ResponseWriter) error {
	return nil
}
