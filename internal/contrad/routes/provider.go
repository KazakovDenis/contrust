package routes

import (
	"log"
	"net/http"
)

func addProvider(w *http.ResponseWriter) error {
	return nil
}

func Provider(w http.ResponseWriter, r *http.Request) {
	var err error

	switch r.Method {
	case http.MethodPost:
		err = addProvider(&w)
	default:
		err = notAllowed(&w)
	}

	if err != nil {
		log.Fatal(err)
	}
}
