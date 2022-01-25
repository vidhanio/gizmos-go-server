package server

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Debug().
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Msg("request")

			next.ServeHTTP(w, r)

			log.Debug().
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Msg("response")
		},
	)
}

func JSONContentTypeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("content-type", "application/json")

			handler.ServeHTTP(w, r)
		},
	)
}
