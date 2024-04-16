package routes

import (
	"net/http"

	"github.com/KazakovDenis/contra/internal/contrad/request"
)

func Index(w http.ResponseWriter, r *http.Request) {
	httpCtx := request.NewHttpContext(&w, r)
	httpCtx.MakeResponse(http.StatusOK, "<h1>Contrad - the central contracts storage</h1>")
}
