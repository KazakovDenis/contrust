package routes

import (
	"net/http"

	"github.com/KazakovDenis/contrust/internal/server/request"
)

func Index(w http.ResponseWriter, r *http.Request) {
	httpCtx := request.NewHttpContext(&w, r)
	httpCtx.MakeResponse(http.StatusOK, "<h1>Contrustd - the centralized contracts storage</h1>", "text/html")
}
