package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aaronland/go-aws-lambda/functionurl"
)

func main() {

	client_uri := flag.String("client-uri", "", "A URI string in the form of 'functionurl://?credentials={CREDENTIALS}&region={AWS_REGION}' where {CREDENTIALS} is a valid aaronland/go-aws-session credentials strings as described in https://github.com/aaronland/go-aws-session#credentials.")
	method := flag.String("method", "GET", "Valid options are: GET, POST.")
	body := flag.String("body", "", "The string body to pass to POST requests. (Eventually this will support more than just strings.)")

	flag.Parse()

	ctx := context.Background()

	cl, err := functionurl.NewClient(ctx, *client_uri)

	if err != nil {
		log.Fatalf("Failed to create new function URL client, %w", err)
	}

	for _, uri := range flag.Args() {

		var rsp *http.Response
		var err error

		switch *method {
		case "GET":
			rsp, err = cl.Get(ctx, uri)
		case "POST":
			rsp, err = cl.Post(ctx, uri, strings.NewReader(*body))
		default:
			log.Fatalf("Invalid method")
		}

		if err != nil {
			log.Fatalf("Failed to GET %s, %v", uri, err)
		}

		defer rsp.Body.Close()

		_, err = io.Copy(os.Stdout, rsp.Body)

		if err != nil {
			log.Fatalf("Failed to read response body from %s, %v", uri, err)
		}
	}
}
