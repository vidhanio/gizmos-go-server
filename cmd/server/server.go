package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	// Configure and run the app

	app := &App{
		Config: config.New(),
	}

	var err error
	app.Database, err = database.New(app.Config, "vidhan-db")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to create database.")
	}

	app.Server = httpd.New(app.Config, app.Database.Database)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to create server.")
	}

	err = app.Server.Start()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start server.")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c

		err := app.Database.Stop()
		if err != nil {
			log.Error().
				Err(err).
				Msg("Failed to stop database.")
		}

		os.Exit(0)
	}()

	app.Database.Database.Client().Disconnect(context.Background())
}
