package scenario

import "github.com/KazakovDenis/contrust/internal/server/request"

type Executable interface {
	Execute(httpCtx *request.HttpContext) error
}

type Scenario struct{}

func (sc *Scenario) Execute(httpCtx *request.HttpContext) (string, error) {
	return "", nil
}
