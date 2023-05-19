package main

import (
	"net/http"
	userC "vominhtrungpro/usermanagement/internal/controller/user"
	userR "vominhtrungpro/usermanagement/internal/handler/rest/user"

	authenC "vominhtrungpro/usermanagement/internal/controller/authentication"
	authenR "vominhtrungpro/usermanagement/internal/handler/rest/authentication"

	"github.com/go-chi/chi/v5"
)

type router struct {
	userCtrl   userC.Controller
	authenCtrl authenC.Controller
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
			r.Post(pattern, userR.New(rtr.userCtrl).CreateUser())
			r.Put(pattern, userR.New(rtr.userCtrl).UpdateUser())
			r.Delete(pattern+"/id/{id}", userR.New(rtr.userCtrl).DeleteUser())
			r.Get(pattern+"/id/{id}", userR.New(rtr.userCtrl).GetUserByID())
		})
		r.Group(func(r chi.Router) {
			pattern := prefix + "/login"

			r.Post(pattern, authenR.New(rtr.authenCtrl).Login())
		})

	})

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{UserCtrl: rtr.userCtrl}}))
	// r.Handle("/graphql", srv)
}
