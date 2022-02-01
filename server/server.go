package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"
	"github.com/vidhanio/gizmos-go-server/db"
)

type GizmoServer struct {
	mux *chi.Mux
	db  db.GizmoDB
	ctx context.Context
}

func New(mux *chi.Mux, db db.GizmoDB) *GizmoServer {
	s := &GizmoServer{
		mux: mux,
		db:  db,
		ctx: context.Background(),
	}

	s.mux.Use(LoggerMiddleware, JSONContentTypeMiddleware, cors.Handler(
		cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		},
	),
	)

	s.mux.Get("/gizmos", s.GetGizmos)
	s.mux.Get("/gizmos/{resource}", s.GetGizmo)
	s.mux.Post("/gizmos", s.PostGizmo)
	s.mux.Put("/gizmos/{resource}", s.PutGizmo)
	s.mux.Delete("/gizmos/{resource}", s.DeleteGizmo)

	return s
}

func (s *GizmoServer) Start() {
	if s.mux == nil {
		log.Fatal().Msg("mux is nil")
	}

	if s.db == nil {
		log.Fatal().Msg("database is nil")
	}

	go func() {
		err := http.ListenAndServe(":8000", s.mux)
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("failed to start server")
		}
	}()
}

func (s *GizmoServer) Stop() error {
	if s.mux == nil {
		return errors.New("mux is nil")
	}

	if s.db == nil {
		return errors.New("database is nil")
	}

	return s.db.Stop()
}

func (s *GizmoServer) ctxTimeout(timeout int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(s.ctx, time.Duration(timeout)*time.Second)
}
