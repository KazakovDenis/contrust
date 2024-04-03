package routes

import (
	"io"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Contrad - the central contracts storage\n")
	if err != nil {
		log.Fatal(err)
	}
}
