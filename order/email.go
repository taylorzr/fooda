package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/pkg/errors"
)

// Order is something
type Order struct {
	Email        string
	DeliveryTime time.Time
}

func parseEmail(data string) (*Order, error) {
	email := AwsEmail{}

	err := json.Unmarshal([]byte(data), &email)

	if err != nil {
		return nil, errors.Wrap(err, "Unable to unmarshal message")
	}

	// fmt.Println("Headers")
	// fmt.Printf("%#v\n", email.Mail.Headers)
	// fmt.Print("------------------------------------------")
	// fmt.Println("Common headers")
	// fmt.Printf("%#v\n", email.Mail.CommonHeaders)
	// fmt.Println("end")

	deliveryTime, err := extractDeliveryTime(email.Content)

	if err != nil {
		return nil, err
	}

	return &Order{
		Email:        email.Mail.Source,
		DeliveryTime: deliveryTime,
	}, nil
}

func extractDeliveryTime(email string) (time.Time, error) {
	r := regexp.MustCompile(`When\s+\w+, (\w{3}) (\d{1,2})\w{2}\s+BETWEEN .*\s+Your Order`)

	match := r.FindStringSubmatch(email)

	if len(match) == 0 {
		return time.Time{}, fmt.Errorf("Unable to match stuff")
	}

	timeString := fmt.Sprintf("%s %s %d 12:00", match[1], match[2], time.Now().Year())

	// Reference time: Mon Jan 2 15:04:05 -0700 MST 2006
	t, err := time.Parse("Jan 2 2006 15:04", timeString)

	if err != nil {
		return time.Time{}, errors.Wrap(err, "Failed to parse order delivery time")
	}

	return t, nil
}
