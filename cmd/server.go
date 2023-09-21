package main

import (
	"go-template/api/database"
	"go-template/api/route"

	"github.com/rs/zerolog/log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to run server")
	}
}

func run() error {
	db, err := database.NewPostgresDatabase()
	if err != nil {
		return err
	}
	err = db.Migrate()
	if err != nil {
		return err
	}
	defer db.Close()

	router := route.NewEchoRouter(db)
	router.RegisterRoutes()

	return router.Start(":1323")
}
