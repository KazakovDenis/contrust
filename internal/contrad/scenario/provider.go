package scenario

import (
	"io"
	"log"
	"net/http"

	"github.com/KazakovDenis/contra/internal/contrad/repo"
)

type AddProviderScenario struct {
	Scenario
	providerName string
}

func (sc *AddProviderScenario) Execute(wr *http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	err := repo.NewProviderRepo(&ctx).Add(sc.providerName)
	_, respErr := io.WriteString(*wr, "OK\n")
	if respErr != nil {
		log.Fatal(err)
	}
	return err
}

func NewAddProviderScenario(name string) *AddProviderScenario {
	return &AddProviderScenario{
		providerName: name,
	}
}
