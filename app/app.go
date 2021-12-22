package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/vidhanio/gizmos-answer-server/app/handlers"
	"github.com/vidhanio/gizmos-answer-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Router   *mux.Router
	Database *mongo.Database
}

func ConfigAndRunApp(config *config.Config) {
	app := new(App)
	app.Initialize(config)
	app.Run(config.ServerHost)
}

func (app *App) Initialize(config *config.Config) {
	log.Info().Msg("Initializing app...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI()))
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to connect to MongoDB.")
	}

	app.Database = client.Database("vidhan-db")
	app.Router = mux.NewRouter()
	app.Router.Use(handlers.JSONContentTypeMiddleware)
	app.setRouters()

	log.Info().Msg("App initialized.")
}

func (app *App) setRouters() {
	app.Get("/gizmos", app.handleRequest(handlers.GetGizmos))
	app.Get("/gizmo/{resource:[0-9]+}", app.handleRequest(handlers.GetGizmo))
	app.Post("/create-gizmo", app.handleRequest(handlers.CreateGizmo))
	app.Put("/update-gizmo/{resource:[0-9]+}", app.handleRequest(handlers.UpdateGizmo))
	app.Delete("/delete-gizmo/{resource:[0-9]+}", app.handleRequest(handlers.DeleteGizmo))
}

func (app *App) Get(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("GET").Queries(queries...)
}
func (app *App) Post(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("POST").Queries(queries...)
}

func (app *App) Put(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("PUT").Queries(queries...)
}

func (app *App) Delete(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("DELETE").Queries(queries...)
}

func (app *App) Run(host string) {
	log.Info().Msg("Starting server...")

	go func() {
		log.Info().Msg("Server started.")

		err := http.ListenAndServe(host, app.Router)
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Failed to start server.")
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Info().Msg("Shutting down server...")

	app.Database.Client().Disconnect(context.Background())

	log.Info().Msg("Server shut down.")
}

func (app *App) handleRequest(handler func(db *mongo.Collection, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.Database.Collection("gizmos"), w, r)
	}
}
