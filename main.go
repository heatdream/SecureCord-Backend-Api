package main

import (
	_ "github.com/lib/pq"

	"ShelterChatBackend/Api/api"
	db "ShelterChatBackend/Api/db/sqlc"
	"ShelterChatBackend/Api/util"
	"database/sql"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewSever(config, &store)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
