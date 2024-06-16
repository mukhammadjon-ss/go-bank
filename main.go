package main

import (
	"context"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(connPool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
}
