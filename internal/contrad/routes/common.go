package routes

import (
	"io"
	"net/http"
)

func notAllowed(w *http.ResponseWriter) error {
	(*w).WriteHeader(http.StatusMethodNotAllowed)
	_, err := io.WriteString(*w, "Not allowed")
	return err
}
