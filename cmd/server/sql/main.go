package main

import (
	"flag"
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

	dbHost := flag.String("db-host", "localhost", "The host of the SQL database")
	dbPort := flag.Int("db-port", 5432, "The port of the SQL database")
	dbUser := flag.String("db-user", "vidhanio", "The user of the SQL database")
	dbPass := flag.String("db-pass", "vidhanio", "The password of the SQL database")
	dbName := flag.String("db-name", "vidhanio", "The name of the SQL database")

	flag.Parse()

	log.Info().
		Msg("Initializing server...")

	server := server.New(sql.New(pgx.ConnConfig{
		Host:     *dbHost,
		Port:     uint16(*dbPort),
		User:     *dbUser,
		Password: *dbPass,
		Database: *dbName,
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
