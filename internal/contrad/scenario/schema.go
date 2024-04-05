package scenario

import "net/http"

type AddSchemaScenario struct {
	Scenario
}

func (sc *AddSchemaScenario) Execute(wr *http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewAddSchemaScenario() *AddSchemaScenario {
	return &AddSchemaScenario{}
}

type GetSchemaScenario struct {
	Scenario
}

func (sc *GetSchemaScenario) Execute(wr *http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewGetSchemaScenario() *GetSchemaScenario {
	return &GetSchemaScenario{}
}
