package main

import (
	"context"
	"os"

	"github.com/jackc/pgx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vidhanio/gizmos-go-server/db/mongodb"
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

	mdb := mongodb.New(client.Database("vidhan-db"))

	log.Info().
		Msg("Connected to database.")

	// use pgx to connect to postgres
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "vidhanio",
		Password: "vidhanio",
		Database: "vidhanio",
	})
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to connect to Postgres.")
	}

	defer conn.Close()

	// insert all gizmos into the postgres database

	gizmos, err := mdb.GetGizmos()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to get gizmos from MongoDB.")
	}

	// create table gizmos

	_, err = conn.Exec("CREATE TABLE gizmos (id SERIAL PRIMARY KEY, title TEXT, materials TEXT, description TEXT, resource INT, answers TEXT[])")
	for _, g := range gizmos {
		_, err := conn.Exec("INSERT INTO gizmos (title, materials, description, resource, answers) VALUES ($1, $2, $3, $4, $5)",
			g.Title, g.Materials, g.Description, g.Resource, g.Answers)
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Failed to insert gizmo into Postgres.")
		}
	}
}
