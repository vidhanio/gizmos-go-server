package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/vidhanio/gizmos-answer-server/config"
	"github.com/vidhanio/gizmos-answer-server/database"
	"github.com/vidhanio/gizmos-answer-server/httpd"
)

type App struct {
	Server   *httpd.Server
	Database *database.Database
	Config   *config.Config
}

func main() {
	// Initialize zerolog

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout})
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	// Configure and run the app

	app := &App{
		Config: config.New(),
	}

	app.Database = database.New(app.Config, "vidhan-db")
	app.Server = httpd.New(app.Config, app.Database.Database)

	go func() {
		app.Server.Start()
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	app.Database.Database.Client().Disconnect(context.Background())
}
