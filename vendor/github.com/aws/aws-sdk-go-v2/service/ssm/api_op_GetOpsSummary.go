// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// View a summary of operations metadata (OpsData) based on specified filters and
// aggregators. OpsData can include information about Amazon Web Services Systems
// Manager OpsCenter operational workitems (OpsItems) as well as information about
// any Amazon Web Services resource or service configured to report OpsData to
// Amazon Web Services Systems Manager Explorer.
func (c *Client) GetOpsSummary(ctx context.Context, params *GetOpsSummaryInput, optFns ...func(*Options)) (*GetOpsSummaryOutput, error) {
	if params == nil {
		params = &GetOpsSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOpsSummary", params, optFns, c.addOperationGetOpsSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOpsSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOpsSummaryInput struct {

	// Optional aggregators that return counts of OpsData based on one or more
	// expressions.
	Aggregators []types.OpsAggregator

	// Optional filters used to scope down the returned OpsData.
	Filters []types.OpsFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	// The OpsData data type to return.
	ResultAttributes []types.OpsResultAttribute

	// Specify the name of a resource data sync to get.
	SyncName *string

	noSmithyDocumentSerde
}

type GetOpsSummaryOutput struct {

	// The list of aggregated details and filtered OpsData.
	Entities []types.OpsEntity

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOpsSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetOpsSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOpsSummary{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetOpsSummary"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpGetOpsSummaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOpsSummary(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// GetOpsSummaryPaginatorOptions is the paginator options for GetOpsSummary
type GetOpsSummaryPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetOpsSummaryPaginator is a paginator for GetOpsSummary
type GetOpsSummaryPaginator struct {
	options   GetOpsSummaryPaginatorOptions
	client    GetOpsSummaryAPIClient
	params    *GetOpsSummaryInput
	nextToken *string
	firstPage bool
}

// NewGetOpsSummaryPaginator returns a new GetOpsSummaryPaginator
func NewGetOpsSummaryPaginator(client GetOpsSummaryAPIClient, params *GetOpsSummaryInput, optFns ...func(*GetOpsSummaryPaginatorOptions)) *GetOpsSummaryPaginator {
	if params == nil {
		params = &GetOpsSummaryInput{}
	}

	options := GetOpsSummaryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetOpsSummaryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetOpsSummaryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetOpsSummary page.
func (p *GetOpsSummaryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetOpsSummaryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetOpsSummary(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// GetOpsSummaryAPIClient is a client that implements the GetOpsSummary operation.
type GetOpsSummaryAPIClient interface {
	GetOpsSummary(context.Context, *GetOpsSummaryInput, ...func(*Options)) (*GetOpsSummaryOutput, error)
}

var _ GetOpsSummaryAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetOpsSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetOpsSummary",
	}
}