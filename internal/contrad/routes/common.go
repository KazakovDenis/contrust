package routes

import (
	"io"
	"net/http"

	"github.com/KazakovDenis/contra/internal/common/log"
)

func NotAllowed(w *http.ResponseWriter) {
	(*w).WriteHeader(http.StatusMethodNotAllowed)
	_, err := io.WriteString(*w, "Not allowed")
	if err != nil {
		log.Error("%s", err)
	}
}

func MakeResponse(w *http.ResponseWriter, status int, response string) {
	(*w).WriteHeader(status)

	if len(response) > 0 {
		_, err := io.WriteString(*w, response)
		if err != nil {
			log.Error("%s", err)
		}
	}
}
