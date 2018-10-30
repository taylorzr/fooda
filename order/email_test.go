package main

import (
	"testing"
)

func TestParseEmail(t *testing.T) {
	data := `{
		"mail": {
			"source": "taylorzr@gmail.com",
			"headers": [
			  {
					"name": "To",
					"value": "some@thing.com"
				}
			],
			"common_headers": {
				"from": [
         "Zach Taylor <taylorzr@gmail.com>"
				],
				"to": [
         "Zach Taylor <taylorzr@gmail.com>"
				]
			}
		},
		"content": "When Saturday, Aug 15th BETWEEN 11:45am-12:15pm Your Order"
	}`

	order, err := parseEmail(data)

	if err != nil {
		t.Error("Expected no error, got", err)
	}

	if order.Email != "taylorzr@gmail.com" {
		t.Error("Expected email of taylorzr@gmail.com, got ", order.Email)
	}

	if order.DeliveryTime.Month() != 8 {
		t.Errorf("Month wrong")
	}

	if order.DeliveryTime.Day() != 15 {
		t.Errorf("Day wrong")
	}

	if order.Email != "taylorzr@gmail.com" {
		t.Error("Expected email of taylorzr@gmail.com, got ", order.Email)
	}
}
