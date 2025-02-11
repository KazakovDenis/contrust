package routes

import (
	"net/http"

	"github.com/KazakovDenis/contrust/internal/common/log"
	"github.com/KazakovDenis/contrust/internal/server/request"
	"github.com/KazakovDenis/contrust/internal/server/scenario"
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
		httpCtx.MakeResponse(http.StatusBadRequest, "Invalid data", "text/plain")
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
	params := httpCtx.Params()
	provider := params.Get("provider")
	if provider == "" {
		httpCtx.MakeResponse(http.StatusBadRequest, "Argument required: provider", "text/plain")
		return
	}

	data, err := scenario.NewGetSchemaScenario().Execute(httpCtx)
	if err != nil {
		httpCtx.MakeResponse(http.StatusInternalServerError, "Internal error", "text/plain")
		return
	}

	payload := map[string]interface{}{
		"result": data,
	}

	httpCtx.MakeJsonResponse(http.StatusOK, payload)
}
