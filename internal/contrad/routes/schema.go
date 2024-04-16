package routes

import (
	"net/http"

	"github.com/KazakovDenis/contra/internal/common/log"
	"github.com/KazakovDenis/contra/internal/contrad/request"
	"github.com/KazakovDenis/contra/internal/contrad/scenario"
)

func Schema(w http.ResponseWriter, r *http.Request) {
	httpCtx := request.NewHttpContext(&w, r)

	switch r.Method {
	case http.MethodGet:
		getSchema(httpCtx)
	case http.MethodPost:
		addSchema(httpCtx)
	default:
		httpCtx.NotAllowed()
	}
}

func addSchema(httpCtx *request.HttpContext) {
	_, err := scenario.NewAddSchemaScenario().Execute(httpCtx)
	if err != nil {
		log.Error("%s", err)
	}
}

func getSchema(httpCtx *request.HttpContext) {
	_, err := scenario.NewGetSchemaScenario().Execute(httpCtx)
	if err != nil {
		log.Error("%s", err)
	}
}
