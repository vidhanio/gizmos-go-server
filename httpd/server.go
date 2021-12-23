package httpd

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/vidhanio/gizmos-answer-server/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Router   *mux.Router
	Database *mongo.Database
	Config   *config.Config
}

func New(cfg *config.Config, db *mongo.Database) *Server {
	s := &Server{
		Database: db,
		Config:   cfg,
	}
	s.Router = mux.NewRouter()
	s.Router.Use(JSONContentTypeMiddleware)
	s.setRoutes()

	return s
}

func (s *Server) Start() error {
	log.Info().Msg("Starting server...")

	err := http.ListenAndServe(s.Config.ServerHost, s.Router)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to start server.")

		return err
	}

	return nil
}

func (s *Server) setRoutes() {
	s.Router.HandleFunc("/gizmos", s.handleRequest(GetGizmos)).Methods("GET")
	s.Router.HandleFunc("/gizmo/{resource:[0-9]+}", s.handleRequest(GetGizmo)).Methods("GET")
	s.Router.HandleFunc("/create-gizmo", s.handleRequest(CreateGizmo)).Methods("POST")
	s.Router.HandleFunc("/update-gizmo/{resource:[0-9]+}", s.handleRequest(UpdateGizmo)).Methods("PUT")
	s.Router.HandleFunc("/delete-gizmo/{resource:[0-9]+}", s.handleRequest(DeleteGizmo)).Methods("DELETE")
}

func (app *Server) handleRequest(handler func(c *mongo.Collection, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.Database.Collection("gizmos"), w, r)
	}
}
