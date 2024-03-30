package v1

import (
	"io"
	"log"
	"net/http"
)

func ApiV1Router(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch ctx.Value(KeyServerAddr) {
	case "/api/v1/":
		Index(w, r)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "This is my website!\n")
	if err != nil {
		log.Fatal(err)
	}
}
