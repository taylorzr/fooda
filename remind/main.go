package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/taylorzr/fooda/fooda"
)

func remind() {
	forgetters, err := fooda.Forgetters()

	if err != nil {
		log.Fatal(err)
	}

	for _, forgetter := range forgetters {
		fooda.Remind(forgetter.Handle)
	}
}

func main() {
	lambda.Start(fooda.Remind)
}
