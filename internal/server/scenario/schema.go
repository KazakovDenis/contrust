package scenario

import (
	"github.com/KazakovDenis/contrust/internal/server/repo"
	"github.com/KazakovDenis/contrust/internal/server/request"
)

type AddSchemaScenario struct {
	Scenario
	provider   string
	contrustct map[string]interface{}
}

func (sc *AddSchemaScenario) Execute(httpCtx *request.HttpContext) (string, error) {
	ctx := httpCtx.Context()
	result, err := repo.NewSchemaRepo(&ctx).Add(sc.provider, sc.contrustct)
	return result, err
}

func NewAddSchemaScenario(provider string, contrustct map[string]interface{}) *AddSchemaScenario {
	return &AddSchemaScenario{
		provider:   provider,
		contrustct: contrustct,
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
