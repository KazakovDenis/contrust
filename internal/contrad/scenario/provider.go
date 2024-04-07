package scenario

import (
	"net/http"

	"github.com/KazakovDenis/contra/internal/contrad/repo"
)

type AddProviderScenario struct {
	Scenario
	providerName string
}

func (sc *AddProviderScenario) Execute(wr *http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	return repo.NewProviderRepo(&ctx).Add(sc.providerName)
}

func NewAddProviderScenario(name string) *AddProviderScenario {
	return &AddProviderScenario{
		providerName: name,
	}
}
