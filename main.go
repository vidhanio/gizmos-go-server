package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/vidhanio/gizmos-answer-server/app"
	"github.com/vidhanio/gizmos-answer-server/config"
)

func main() {
	// Initialize zerolog

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout})
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	// Configure and run the app

	config := config.New()
	app.ConfigAndRunApp(config)
}
