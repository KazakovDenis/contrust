package scenario

import (
	"github.com/KazakovDenis/contra/internal/contrad/http"
	"github.com/KazakovDenis/contra/internal/contrad/repo"
)

type AddProviderScenario struct {
	Scenario
	providerName string
}

func (sc *AddProviderScenario) Execute(httpCtx *http.HttpContext) (string, error) {
	ctx := httpCtx.Context()
	result, err := repo.NewProviderRepo(&ctx).Add(sc.providerName)
	return result, err
}

func NewAddProviderScenario(name string) *AddProviderScenario {
	return &AddProviderScenario{
		providerName: name,
	}
}
