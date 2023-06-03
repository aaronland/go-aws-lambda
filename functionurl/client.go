package functionurl

// https://docs.aws.amazon.com/lambda/latest/dg/urls-invocation.html
// https://docs.aws.amazon.com/IAM/latest/UserGuide/create-signed-request.html
// https://docs.aws.amazon.com/sdk-for-go/api/aws/signer/v4/

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	aa_session "github.com/aaronland/go-aws-session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
)

// Client is a struct that implements methods for signing and executing requests to AWS Lambda Function URLs.
type Client struct {
	config      *aws.Config
	client      *http.Client
	credentials string
	region      string
	ttl         time.Duration
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

	http_client := &http.Client{}

	ttl := 1 * time.Minute

	cl := &Client{
		region:      q_region,
		client:      http_client,
		credentials: q_credentials,
		config:      cfg,
		ttl:         ttl,
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

// SignRequest will sign 'req' and 'body' and apply the necessary AWS Signature request headers
// to 'req'.
func (cl *Client) SignRequest(ctx context.Context, req *http.Request, body io.ReadSeeker) error {

	switch cl.credentials {
	case aa_session.AnonymousCredentialsString, aa_session.IAMCredentialsString:
		return nil
	default:
		// carry on
	}

	creds := cl.config.Credentials
	signer := v4.NewSigner(creds)

	service := "lambda"
	now := time.Now()

	_, err := signer.Presign(req, body, service, cl.region, cl.ttl, now)

	if err != nil {
		return fmt.Errorf("Failed to sign request, %w", err)
	}

	return nil
}
