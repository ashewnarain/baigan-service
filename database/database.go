package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "baigan.c4ygd1qu2m2b.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "baiganadmin"
	password = "ba!gan2018"
	dbname   = "baigandb"
)

var SQLDB *sql.DB

// ConnectDB opens a connection to the database
func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	SQLDB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = SQLDB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}
