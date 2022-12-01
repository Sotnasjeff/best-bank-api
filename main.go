package main

import (
	"database/sql"
	"log"

	"github.com/best-bank-api/util"

	"github.com/best-bank-api/api"
	db "github.com/best-bank-api/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Couldn't load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Couldn't connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Can't create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can't start server:", err)
	}
}
