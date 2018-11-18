package fooda

import (
	"reflect"
	"testing"
	"time"
)

func TestTodaysOrders(t *testing.T) {
	// order := Order{
	// 	Restaurant:   "Taco Bell",
	// 	DeliveryDate: time.Now(),
	// 	UserID:       1,
	// }

	// _, err := db.NamedExec("insert into orders(user_id, delivery_date, restaurant) values (:user_id, :delivery_date, :restaurant)", order)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	orders, err := TodaysOrders()

	if err != nil {
		t.Error(err)
	}

	actual := make([]string, len(orders))

	for i, order := range orders {
		actual[i] = order.Restaurant
	}

	expected := []string{"Burger King"}

	ok := reflect.DeepEqual(actual, expected)

	if !ok {
		t.Errorf("Expected restaurants to match\n%s\n%s", expected, actual)
	}
}

func TestCreateOrder(t *testing.T) {
	err := CreateOrder(
		"taylorzr@gmail.com",
		time.Date(2018, 11, 15, 0, 0, 0, 0, time.UTC),
		"Taco Bell",
	)

	if err != nil {
		t.Error("Expected no error, got", err)
	}
}
