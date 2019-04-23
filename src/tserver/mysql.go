package tserver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	MySQL *sql.DB
)

// fixme not a pool
func init() {
	// fixme hardcore username password
	db, err := sql.Open("mysql", "root:Tomato-123@/student")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	MySQL = db
}
