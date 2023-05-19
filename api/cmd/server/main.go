package main

import (
	"context"
	"database/sql"
	"log"
	"vominhtrungpro/usermanagement/internal/appconfig/db/pg"
	"vominhtrungpro/usermanagement/internal/controller/authentication"
	"vominhtrungpro/usermanagement/internal/controller/user"
	"vominhtrungpro/usermanagement/internal/httpserver"
	"vominhtrungpro/usermanagement/internal/repository"
	"vominhtrungpro/usermanagement/internal/repository/generator"
)

func main() {
	ctx := context.Background()

	conn, err := pg.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	generator.InitSnowflakeGenerators()

	rtr, err := initRouter(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	handler := httpserver.Handler(ctx, rtr.routes)
	httpserver.Start(handler)
}

func initRouter(_ context.Context, db *sql.DB) (router, error) {
	repo := repository.New(db)

	userCtrl := user.New(repo)

	authenCtrl := authentication.New(repo)

	return router{
		userCtrl:   userCtrl,
		authenCtrl: authenCtrl,
	}, nil
}
