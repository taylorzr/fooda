package main

import (
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	err := createOrder(
		"taylorzr@gmail.com",
		time.Date(2018, 11, 15, 0, 0, 0, 0, time.UTC),
		"Taco Bell",
	)

	if err != nil {
		t.Error("Expected no error, got", err)
	}
}
