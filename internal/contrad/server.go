package contrad

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"

	"github.com/KazakovDenis/contra/internal/common/log"
	"github.com/KazakovDenis/contra/internal/contrad/constants"
	"github.com/KazakovDenis/contra/internal/contrad/mongodb"
	"github.com/KazakovDenis/contra/internal/contrad/routes"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.Index)
	mux.HandleFunc("/provider", routes.ProviderRouter)
	mux.HandleFunc("/schema", routes.SchemaRouter)
	return mux
}

func newServer(cfg *AppConfig) (*http.Server, *context.Context, func()) {
	log.InitLogger(Config.LogLevel, Config.LogFormat)

	ctx := context.Background()
	db, disconnect := mongodb.Connect(ctx, Config.DatabaseURI, Config.DatabaseName)

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: newMux(),
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, constants.KeyServerAddr, l.Addr().String())
			ctx = context.WithValue(ctx, constants.Database, db)
			return ctx
		},
	}
	return srv, &ctx, disconnect
}

func Run() {
	log.Info("Contrad is running on http://0.0.0.0:%s", Config.ServerPort)

	srv, ctx, shutdown := newServer(Config)
	defer shutdown()

	err := srv.ListenAndServe()
	<-(*ctx).Done()

	if errors.Is(err, http.ErrServerClosed) {
		log.Error("%s\n", err)
	} else if err != nil {
		log.Error("Error while starting the server: %s\n", err)
		os.Exit(1)
	}
}
