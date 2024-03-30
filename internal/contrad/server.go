package contrad

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	v1 "github.com/KazakovDenis/contra/internal/contrad/v1"
)

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/api/v1/", v1.ApiV1Router)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	bind := fmt.Sprintf(":%s", port)
	log.Printf("Start serving on %s", bind)

	ctx := context.Background()
	srv := &http.Server{
		Addr:    bind,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, v1.KeyServerAddr, l.Addr().String())
			return ctx
		},
	}
	err := srv.ListenAndServe()
	<-ctx.Done()

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("%s\n", err)
	} else if err != nil {
		log.Printf("Error while starting the server: %s\n", err)
		os.Exit(1)
	}
}
