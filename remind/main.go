package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/taylorzr/fooda/fooda"
)

func main() {
	lambda.Start(fooda.Remind)
}
