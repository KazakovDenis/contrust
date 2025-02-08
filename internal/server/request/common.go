package request

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/KazakovDenis/contrust/internal/common/log"
)

type HttpContext struct {
	writer  *http.ResponseWriter
	request *http.Request
}

func NewHttpContext(writer *http.ResponseWriter, request *http.Request) *HttpContext {
	return &HttpContext{
		writer:  writer,
		request: request,
	}
}

func (httpCtx *HttpContext) Context() context.Context {
	return httpCtx.request.Context()
}

func (httpCtx *HttpContext) Json() (map[string]interface{}, error) {
	var jsonData map[string]interface{}
	err := json.NewDecoder(httpCtx.request.Body).Decode(&jsonData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (httpCtx *HttpContext) Params() url.Values {
	return httpCtx.request.URL.Query()
}

func (httpCtx *HttpContext) MakeResponse(status int, response, contentType string) {
	w := httpCtx.writer
	(*w).WriteHeader(status)

	if len(response) > 0 {
		(*w).Header().Set("Content-Type", contentType)
		_, err := io.WriteString(*w, response)
		if err != nil {
			log.Error("%s", err)
		}
	}
}

func (httpCtx *HttpContext) MakeJsonResponse(status int, response interface{}) {
	w := httpCtx.writer
	(*w).WriteHeader(status)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Error("%s", err)
		http.Error(*w, "Internal error", http.StatusInternalServerError)
		return
	}

	(*w).Header().Set("Content-Type", "application/json")

	_, err = io.WriteString(*w, string(jsonResponse))
	if err != nil {
		log.Error("%s", err)
	}
}

func (httpCtx *HttpContext) NotAllowed() {
	httpCtx.MakeResponse(http.StatusMethodNotAllowed, "Not allowed", "text/plain")
}
