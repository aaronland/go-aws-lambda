package main

import (
	"context"
	"flag"
	"log"
	"strings"

	"github.com/aaronland/go-aws-lambda"
)

func main() {

	lambda_uri := flag.String("lambda-uri", "", "...")
	json := flag.Bool("json", false, "...")

	flag.Parse()

	ctx := context.Background()

	f, err := lambda.NewLambdaFunction(ctx, *lambda_uri)

	if err != nil {
		log.Fatalf("Failed to create new Lambda function, %v", err)
	}

	payload := strings.Join(flag.Args(), " ")

	if *json {
		_, err = f.InvokeWithJSON(ctx, []byte(payload))
	} else {
		_, err = f.Invoke(ctx, payload)
	}

	if err != nil {
		log.Fatalf("Failed to invoke function, %v", err)
	}
}
