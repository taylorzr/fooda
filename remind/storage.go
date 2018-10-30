package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Order
type Order struct {
	ID         int
	Restaurant string
	Email      string
	Handle     string
}

// TODO: Return slice of orders
func getOrders() Order {
	db, err := sql.Open("postgres", "dbname=fooda sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	order := Order{}

	query := `
	  SELECT orders.id, orders.restaurant, users.Email, users.handle FROM orders
		JOIN users on users.id = orders.user_id
		WHERE orders.delivery_date = current_date
	`

	err = db.QueryRow(query).Scan(&order.ID, &order.Restaurant, &order.Email, &order.Handle)

	if err != nil {
		log.Fatal(err)
	}

	return order
}
