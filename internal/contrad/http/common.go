package http

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/KazakovDenis/contra/internal/common/log"
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

func (httpCtx *HttpContext) MakeResponse(status int, response string) {
	w := httpCtx.writer
	(*w).WriteHeader(status)

	if len(response) > 0 {
		_, err := io.WriteString(*w, response)
		if err != nil {
			log.Error("%s", err)
		}
	}
}

func NotAllowed(w *http.ResponseWriter) {
	(*w).WriteHeader(http.StatusMethodNotAllowed)
	_, err := io.WriteString(*w, "Not allowed")
	if err != nil {
		log.Error("%s", err)
	}
}
