package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vidhanio/gizmos-go-server/db/mongodb"
	"github.com/vidhanio/gizmos-go-server/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().
		Msg("Connecting to database.")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal().
			Msg("Failed to create MongoDB client.")
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal().
			Msg("Failed to connect to MongoDB.")
	}

	database := client.Database("vidhan-db")

	log.Info().
		Msg("Connected to database.")

	mux := chi.NewRouter()

	server := server.New(mux, mongodb.New(database))

	log.Info().
		Msg("Starting server...")

	server.Start()

	log.Info().
		Msg("Server started.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sc
	log.Info().
		Msg("Stopping server...")

	err = server.Stop()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to stop server.")
	}

	log.Info().
		Msg("Server stopped.")
}
