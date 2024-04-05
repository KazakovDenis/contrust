package scenario

import "net/http"

type Executable interface {
	Execute(wr *http.ResponseWriter, r *http.Request) error
}

type Scenario struct{}

func (sc *Scenario) Execute(wr *http.ResponseWriter, r *http.Request) error {
	return nil
}
