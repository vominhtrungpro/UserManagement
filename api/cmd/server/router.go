package main

import (
	"net/http"
	userC "vominhtrungpro/usermanagement/internal/controller/user"
	userR "vominhtrungpro/usermanagement/internal/handler/rest/user"

	"github.com/go-chi/chi/v5"
)

type router struct {
	userCtrl userC.Controller
}

func (rtr router) routes(r chi.Router) {
	r.Group(rtr.public)
}

func (rtr router) public(r chi.Router) {
	prefix := "/api/public"

	r.Group(func(r chi.Router) {
		r.Get(prefix+"/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world!!"))
		})
	})

	// v1
	r.Group(func(r chi.Router) {
		prefix = prefix + "/v1"

		//User
		r.Group(func(r chi.Router) {
			pattern := prefix + "/users"

			r.Get(pattern, userR.New(rtr.userCtrl).GetAllUser())
		})
	})

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{UserCtrl: rtr.userCtrl}}))
	// r.Handle("/graphql", srv)
}
