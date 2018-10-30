package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "dbname=fooda sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
}

// User is something
type User struct {
	ID     int
	Email  string
	Handle string
}

func createOrder(email string, deliveryDate time.Time, restaurant string) error {
	user := User{}

	query := `
	  SELECT users.id, users.email, users.handle FROM users
		WHERE users.email = $1
	`

	err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Handle)

	if err != nil {
		return err
	}

	insertQuery := `
	  INSERT INTO orders(user_id, delivery_date, restaurant)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id int
	err = db.QueryRow(insertQuery, user.ID, "2014-11-15", "McDonalds").Scan(&id)

	if err != nil {
		return err
	}

	log.Println("New order with id", id)

	return nil
}
