package main

import (
	"context"
	"flag"
	"github.com/aaronland/go-aws-lambda"
	"log"
	"strings"
)

func main() {

	lambda_uri := flag.String("lambda-uri", "", "...")

	flag.Parse()

	ctx := context.Background()

	f, err := lambda.NewLambdaFunction(ctx, *lambda_uri)

	if err != nil {
		log.Fatalf("Failed to create new Lambda function, %v", err)
	}

	payload := strings.Join(flag.Args(), " ")

	_, err = f.Invoke(ctx, payload)

	if err != nil {
		log.Fatalf("Failed to invoke function, %v", err)
	}
}
