package scenario

import (
	"net/http"

	"github.com/KazakovDenis/contra/internal/contrad/repo"
)

type AddProviderScenario struct {
	Scenario
	providerName string
}

func (sc *AddProviderScenario) Execute(wr *http.ResponseWriter, r *http.Request) (string, error) {
	ctx := r.Context()
	result, err := repo.NewProviderRepo(&ctx).Add(sc.providerName)
	return result, err
}

func NewAddProviderScenario(name string) *AddProviderScenario {
	return &AddProviderScenario{
		providerName: name,
	}
}
