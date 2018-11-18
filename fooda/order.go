package fooda

import (
	"log"
	"time"
)

func CreateOrder(email string, deliveryDate time.Time, restaurant string) error {
	// user := User{}

	// query := `
	//   SELECT users.id, users.email, users.handle FROM users
	// 	WHERE users.email = $1
	// `

	// err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Handle)
	user, err := GetUser(email)

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

func TodaysOrders() ([]Order, error) {
	query := `
	  SELECT orders.id, orders.restaurant, users.Email, users.handle FROM orders
		JOIN users on users.id = orders.user_id
		WHERE orders.delivery_date = current_date
	`

	rows, err := db.Queryx(query)

	if err != nil {
		return nil, err
	}

	orders := []Order{}

	for rows.Next() {
		order := Order{}
		err := rows.StructScan(&order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

type Order struct {
	ID           int
	UserID       int `db:"user_id"`
	Restaurant   string
	DeliveryDate time.Time `db:"delivery_date"`
	User
}
