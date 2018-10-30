package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handleSNS)
}

func handleSNS(snsEvent events.SNSEvent) error {
	log.Printf("Processing %d events\n", len(snsEvent.Records))

	for _, record := range snsEvent.Records {
		_, err := parseEmail(record.SNS.Message)

		if err != nil {
			return err
		}
	}

	return nil
}
