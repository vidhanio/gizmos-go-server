package database

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/vidhanio/gizmos-answer-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Database *mongo.Database
	Config   *config.Config
}

func New(cfg *config.Config, name string) *Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info().Msg("Connecting to database...")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI()))
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to connect to MongoDB.")
	}

	return &Database{
		Database: client.Database(name),
		Config:   cfg,
	}
}
