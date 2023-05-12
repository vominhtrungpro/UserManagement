package httpserver

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Handler returns http.Handler
func Handler(ctx context.Context, routerFn func(r chi.Router)) http.Handler {
	r := chi.NewRouter()
	r.Group(routerFn)

	return r
}
