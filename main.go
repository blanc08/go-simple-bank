package main

import (
	"database/sql"
	"log"

	"github.com/blanc08/go-simple-bank/api"
	db "github.com/blanc08/go-simple-bank/db/sqlc"
	"github.com/blanc08/go-simple-bank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the database", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create the server :", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server :", err)
	}

}
