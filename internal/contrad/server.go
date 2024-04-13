package contrad

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/KazakovDenis/contra/internal/contrad/contants"
	"github.com/KazakovDenis/contra/internal/contrad/database"
	"github.com/KazakovDenis/contra/internal/contrad/routes"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.Index)
	mux.HandleFunc("/provider", routes.Provider)
	mux.HandleFunc("/schema", routes.Schema)
	return mux
}

func newServer(cfg *AppConfig) (*http.Server, *context.Context, func()) {
	ctx := context.Background()
	db, disconnect := database.Connect(ctx, Config.DatabaseURI, Config.DatabaseName)

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: newMux(),
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, contants.KeyServerAddr, l.Addr().String())
			ctx = context.WithValue(ctx, contants.Database, db)
			return ctx
		},
	}
	return srv, &ctx, disconnect
}

func Run() {
	log.Printf("Contrad is running on http://0.0.0.0:%s", Config.ServerPort)

	srv, ctx, shutdown := newServer(Config)
	defer shutdown()

	err := srv.ListenAndServe()
	<-(*ctx).Done()

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("%s\n", err)
	} else if err != nil {
		log.Printf("Error while starting the server: %s\n", err)
		os.Exit(1)
	}
}
