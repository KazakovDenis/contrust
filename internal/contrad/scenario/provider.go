package scenario

import "net/http"

type AddProviderScenario struct {
	Scenario
}

func (sc *AddProviderScenario) Execute(wr *http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewAddProviderScenario() *AddProviderScenario {
	return &AddProviderScenario{}
}
