package scenario

import "github.com/KazakovDenis/contra/internal/contrad/request"

type AddSchemaScenario struct {
	Scenario
}

func (sc *AddSchemaScenario) Execute(httpCtx *request.HttpContext) (string, error) {
	return "", nil
}

func NewAddSchemaScenario() *AddSchemaScenario {
	return &AddSchemaScenario{}
}

type GetSchemaScenario struct {
	Scenario
}

func (sc *GetSchemaScenario) Execute(httpCtx *request.HttpContext) (string, error) {
	return "", nil
}

func NewGetSchemaScenario() *GetSchemaScenario {
	return &GetSchemaScenario{}
}
