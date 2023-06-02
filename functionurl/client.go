package functionurl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	aa_session "github.com/aaronland/go-aws-session"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
)

// Client is a struct that implements methods for signing and executing requests to AWS Lambda Function URLs.
type Client struct {
	signer *v4.Signer
	client *http.Client
	region string
}

// NewClient returns a new `Client` instance which is used for signing and executing requests to AWS Lambda Function
// URLs, configured by 'uri' which is expected to take the form of:
//
//	functionurl://?credentials={CREDENTIALS}&region={AWS_REGION}
//
// Where {CREDENTIALS} is a valid aaronland/go-aws-session credentials strings as described in https://github.com/aaronland/go-aws-session#credentials
func NewClient(ctx context.Context, uri string) (*Client, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	q_credentials := q.Get("credentials")
	q_region := q.Get("region")

	cfg, err := aa_session.NewConfigWithCredentials(q_credentials)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse credentials, %w", err)
	}

	creds := cfg.Credentials

	signer := v4.NewSigner(creds)

	http_client := &http.Client{}

	cl := &Client{
		signer: signer,
		region: q_region,
		client: http_client,
	}

	return cl, nil
}

// Get will sign and execute an HTTP GET request to 'uri'.
func (cl *Client) Get(ctx context.Context, uri string) (*http.Response, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)

	if err != nil {
		return nil, fmt.Errorf("Failed to create request, %w", err)
	}

	err = cl.SignRequest(ctx, req, nil)

	if err != nil {
		return nil, fmt.Errorf("Failed to sign request, %w", err)
	}

	return cl.client.Do(req)
}

// Get will sign and execute an HTTP POST request to 'uri' passing 'body' as the request body.
func (cl *Client) Post(ctx context.Context, uri string, body io.ReadSeeker) (*http.Response, error) {

	req, err := http.NewRequestWithContext(ctx, "POST", uri, body)

	if err != nil {
		return nil, fmt.Errorf("Failed to create request, %w", err)
	}

	err = cl.SignRequest(ctx, req, body)

	if err != nil {
		return nil, fmt.Errorf("Failed to sign request, %w", err)
	}

	return cl.client.Do(req)
}

// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/aws/signer/v4#Signer

// SignRequest will sign 'req' and 'body' and apply the necessary AWS Signature request headers
// to 'req'.
func (cl *Client) SignRequest(ctx context.Context, req *http.Request, body io.ReadSeeker) error {

	service := "lambda"
	now := time.Now()

	expires := 1 * time.Minute

	_, err := cl.signer.Presign(req, body, service, cl.region, expires, now)

	if err != nil {
		return fmt.Errorf("Failed to sign request, %w", err)
	}

	return nil
}
