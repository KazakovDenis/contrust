package contrad

import (
	"github.com/KazakovDenis/contra/internal/contrad/v1"
	"io"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Printf("%s: got / request\n", ctx.Value(v1.KeyServerAddr))
	_, err := io.WriteString(w, "Index!\n")
	if err != nil {
		log.Fatal(err)
	}
}
