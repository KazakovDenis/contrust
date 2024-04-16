package scenario

import "github.com/KazakovDenis/contra/internal/contrad/request"

type Executable interface {
	Execute(httpCtx *request.HttpContext) error
}

type Scenario struct{}

func (sc *Scenario) Execute(httpCtx *request.HttpContext) (string, error) {
	return "", nil
}
