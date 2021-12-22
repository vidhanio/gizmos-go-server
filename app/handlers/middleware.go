package handlers

import "net/http"

func JSONContentTypeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("content-type", "application/json")
			handler.ServeHTTP(w, r)
		},
	)
}
