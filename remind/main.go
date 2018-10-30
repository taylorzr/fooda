package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

// TODO
func remind(ctx context.Context) {
	// get all orders
	// get all users
	// for each user without order, remind them
}

func main() {
	lambda.Start(remind)
}
