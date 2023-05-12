package main

import (
	"context"
	"log"
	"vominhtrungpro/usermanagement/internal/appconfig/db/pg"
)

func main() {
	ctx := context.Background()

	conn, err := pg.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}
