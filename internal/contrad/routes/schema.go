package routes

import (
	"net/http"

	"github.com/KazakovDenis/contra/internal/common/log"
	"github.com/KazakovDenis/contra/internal/contrad/request"
	"github.com/KazakovDenis/contra/internal/contrad/scenario"
)

func SchemaRouter(w http.ResponseWriter, r *http.Request) {
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
	payload, err := httpCtx.Json()
	if err != nil {
		httpCtx.MakeResponse(http.StatusBadRequest, "Invalid data")
		return
	}

	if err = addSchemaValidate(payload); err != nil {
		log.Error("%s", err)
		return
	}

	_, err = scenario.NewAddSchemaScenario(
		payload["provider"].(string),
		payload["contract"].(map[string]interface{}),
	).Execute(httpCtx)

	if err != nil {
		log.Error("%s", err)
	}
}

func addSchemaValidate(payload map[string]interface{}) error {
	return nil
}

func getSchema(httpCtx *request.HttpContext) {
	_, err := scenario.NewGetSchemaScenario().Execute(httpCtx)
	if err != nil {
		log.Error("%s", err)
	}
}
