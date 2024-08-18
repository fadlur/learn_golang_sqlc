package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/fadlur/learn_golang_sqlc/util"
	_ "github.com/lib/pq"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://postgres:147789@localhost:5432/simple_bank?sslmode=disable"
// )

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	// var err error
	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
