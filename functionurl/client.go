package functionurl

// https://docs.aws.amazon.com/lambda/latest/dg/urls-invocation.html
// https://docs.aws.amazon.com/IAM/latest/UserGuide/create-signed-request.html
// https://docs.aws.amazon.com/sdk-for-go/api/aws/signer/v4/

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/aaronland/go-aws-auth"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Client is a struct that implements methods for signing and executing requests to AWS Lambda Function URLs.
type Client struct {
	config      aws.Config
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

	cfg_uri := fmt.Sprintf("aws://%s?credentials=%s", q_region, q_credentials)
	cfg, err := auth.NewConfig(ctx, cfg_uri)

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
	case auth.AnonymousCredentialsString, auth.IAMCredentialsString:
		return nil
	default:
		// carry on
	}

	creds, err := cl.config.Credentials.Retrieve(ctx)

	if err != nil {
		return err
	}

	h := sha256.New()
	_, err = io.Copy(h, body)

	if err != nil {
		return fmt.Errorf("Failed to copy body to hash, %w", err)
	}

	payload_hash := fmt.Sprintf("%x", h.Sum(nil))

	signer := v4.NewSigner()

	service := "lambda"
	now := time.Now()

	expires := cl.ttl
	query := req.URL.Query()
	query.Set("X-Amz-Expires", strconv.FormatInt(int64(expires/time.Second), 10))
	req.URL.RawQuery = query.Encode()

	_, _, err = signer.PresignHTTP(ctx, creds, req, payload_hash, service, cl.region, now)

	if err != nil {
		return fmt.Errorf("Failed to sign request, %w", err)
	}

	return nil
}
