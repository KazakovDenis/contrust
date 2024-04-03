package contrad

import (
	"context"
	"errors"
	"github.com/KazakovDenis/contra/internal/contrad/routes"
	"log"
	"net"
	"net/http"
	"os"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.Index)
	mux.HandleFunc("/provider", routes.Provider)
	mux.HandleFunc("/schema", routes.Schema)
	return mux
}

func newServer(cfg *Config) (*http.Server, *context.Context) {
	ctx := context.Background()
	srv := &http.Server{
		Addr:    ":" + cfg.serverPort,
		Handler: newMux(),
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, KeyServerAddr, l.Addr().String())
			return ctx
		},
	}
	return srv, &ctx
}

func Run() {
	config := NewConfig()
	log.Printf("Contrad is running on http://0.0.0.0:%s", config.serverPort)

	srv, ctx := newServer(config)
	err := srv.ListenAndServe()
	<-(*ctx).Done()

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("%s\n", err)
	} else if err != nil {
		log.Printf("Error while starting the server: %s\n", err)
		os.Exit(1)
	}
}
