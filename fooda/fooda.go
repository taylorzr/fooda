package fooda

import (
	"context"
	"log"
)

func Remind(ctx context.Context) {
	orders, err := TodaysOrders()

	if err != nil {
		log.Fatal(err)
	}

	users, err := AllUsers()

	missing := []User{}

User:
	for _, user := range users {
		for _, order := range orders {
			if user.Email == order.User.Email {
				break User
			}
		}

		missing = append(missing, user)
	}

}
