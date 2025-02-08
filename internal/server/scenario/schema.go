package scenario

import (
	"github.com/KazakovDenis/contrust/internal/server/repo"
	"github.com/KazakovDenis/contrust/internal/server/request"
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
	provider  string
	transport string
}

func (sc *GetSchemaScenario) Execute(httpCtx *request.HttpContext) ([]repo.Document, error) {
	ctx := httpCtx.Context()
	result, err := repo.NewSchemaRepo(&ctx).Get(sc.provider, &sc.transport)
	return result, err
}

func NewGetSchemaScenario() *GetSchemaScenario {
	return &GetSchemaScenario{}
}
