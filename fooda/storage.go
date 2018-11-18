package fooda

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func init() {
	var err error

	db, err = sqlx.Connect("postgres", "dbname=fooda sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
}
