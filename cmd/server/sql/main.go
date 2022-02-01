package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vidhanio/gizmos-go-server/db/sql"
	"github.com/vidhanio/gizmos-go-server/server"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().
		Msg("Initializing server...")

	server := server.New(sql.New(pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "vidhanio",
		Password: "vidhanio",
		Database: "vidhanio",
	}))

	log.Info().
		Msg("Server initialized.")

	log.Info().
		Msg("Starting server...")

	go func() {
		err := server.Start()
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Failed to start server.")
		}
	}()

	log.Info().
		Msg("Server started.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sc
	log.Info().
		Msg("Stopping server...")

	err := server.Stop()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to stop server.")
	}

	log.Info().
		Msg("Server stopped.")
}
