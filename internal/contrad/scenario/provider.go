package scenario

import (
	"github.com/KazakovDenis/contra/internal/contrad/repo"
	"github.com/KazakovDenis/contra/internal/contrad/request"
)

type AddProviderScenario struct {
	Scenario
	providerName string
}

func (sc *AddProviderScenario) Execute(httpCtx *request.HttpContext) (string, error) {
	ctx := httpCtx.Context()
	result, err := repo.NewProviderRepo(&ctx).Add(sc.providerName)
	return result, err
}

func NewAddProviderScenario(name string) *AddProviderScenario {
	return &AddProviderScenario{
		providerName: name,
	}
}
