package scenario

import (
	"github.com/KazakovDenis/contra/internal/contrad/repo"
	"github.com/KazakovDenis/contra/internal/contrad/request"
)

type AddSchemaScenario struct {
	Scenario
	provider string
	contract map[string]interface{}
}

func (sc *AddSchemaScenario) Execute(httpCtx *request.HttpContext) (string, error) {
	ctx := httpCtx.Context()
	result, err := repo.NewSchemaRepo(&ctx).Add(sc.provider, sc.contract)
	return result, err
}

func NewAddSchemaScenario(provider string, contract map[string]interface{}) *AddSchemaScenario {
	return &AddSchemaScenario{
		provider: provider,
		contract: contract,
	}
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
