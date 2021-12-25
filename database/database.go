package database

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/vidhanio/gizmos-go-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Database *mongo.Database
	Config   *config.Config
}

func New(cfg *config.Config, name string) (*Database, error) {
	log.Info().Msg("Connecting to MongoDB...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI()))
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to connect to MongoDB.")

		return nil, err
	}

	log.Info().Msg("Connected to MongoDB.")

	return &Database{
		Database: client.Database(name),
		Config:   cfg,
	}, nil
}

func (db *Database) Stop() error {
	log.Info().Msg("Disconnecting from MongoDB...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := db.Database.Client().Disconnect(ctx)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to disconnect from MongoDB.")

		return err
	}

	return nil
}
