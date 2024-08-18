package main

import (
	"database/sql"
	"log"

	"github.com/fadlur/learn_golang_sqlc/api"
	db "github.com/fadlur/learn_golang_sqlc/db/sqlc"
	"github.com/fadlur/learn_golang_sqlc/util"
	_ "github.com/lib/pq"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://postgres:147789@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "127.0.0.1:8080"
// )

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("can not start server")
	}
}
