// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the managed policies that are available in your Amazon Web Services
// account, including your own customer-defined managed policies and all Amazon Web
// Services managed policies.
//
// You can filter the list of policies that is returned using the optional
// OnlyAttached , Scope , and PathPrefix parameters. For example, to list only the
// customer managed policies in your Amazon Web Services account, set Scope to
// Local . To list only Amazon Web Services managed policies, set Scope to AWS .
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// For more information about managed policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. For example, this operation does not return tags, even though they
// are an attribute of the returned object. To view all of the information for a
// customer manged policy, see [GetPolicy].
//
// [GetPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func (c *Client) ListPolicies(ctx context.Context, params *ListPoliciesInput, optFns ...func(*Options)) (*ListPoliciesOutput, error) {
	if params == nil {
		params = &ListPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPolicies", params, optFns, c.addOperationListPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPoliciesInput struct {

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true .
	//
	// If you do not include this parameter, the number of items defaults to 100. Note
	// that IAM might return fewer results, even when there are more results available.
	// In that case, the IsTruncated response element returns true , and Marker
	// contains a value to include in the subsequent call that tells the service where
	// to continue from.
	MaxItems *int32

	// A flag to filter the results to only the attached policies.
	//
	// When OnlyAttached is true , the returned list contains only the policies that
	// are attached to an IAM user, group, or role. When OnlyAttached is false , or
	// when the parameter is not included, all policies are returned.
	OnlyAttached bool

	// The path prefix for filtering the results. This parameter is optional. If it is
	// not included, it defaults to a slash (/), listing all policies. This parameter
	// allows (through its [regex pattern]) a string of characters consisting of either a forward
	// slash (/) by itself or a string that must begin and end with forward slashes. In
	// addition, it can contain any ASCII character from the ! ( \u0021 ) through the
	// DEL character ( \u007F ), including most punctuation characters, digits, and
	// upper and lowercased letters.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	PathPrefix *string

	// The policy usage method to use for filtering the results.
	//
	// To list only permissions policies, set PolicyUsageFilter to PermissionsPolicy .
	// To list only the policies used to set permissions boundaries, set the value to
	// PermissionsBoundary .
	//
	// This parameter is optional. If it is not included, all policies are returned.
	PolicyUsageFilter types.PolicyUsageType

	// The scope to use for filtering the results.
	//
	// To list only Amazon Web Services managed policies, set Scope to AWS . To list
	// only the customer managed policies in your Amazon Web Services account, set
	// Scope to Local .
	//
	// This parameter is optional. If it is not included, or if it is set to All , all
	// policies are returned.
	Scope types.PolicyScopeType

	noSmithyDocumentSerde
}

// Contains the response to a successful [ListPolicies] request.
//
// [ListPolicies]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListPolicies.html
type ListPoliciesOutput struct {

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated bool

	// When IsTruncated is true , this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// A list of policies.
	Policies []types.Policy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPolicies"); err != nil {
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
	if err = addSpanRetryLoop(stack, options); err != nil {
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
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicies(options.Region), middleware.Before); err != nil {
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
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

// ListPoliciesPaginatorOptions is the paginator options for ListPolicies
type ListPoliciesPaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true .
	//
	// If you do not include this parameter, the number of items defaults to 100. Note
	// that IAM might return fewer results, even when there are more results available.
	// In that case, the IsTruncated response element returns true , and Marker
	// contains a value to include in the subsequent call that tells the service where
	// to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPoliciesPaginator is a paginator for ListPolicies
type ListPoliciesPaginator struct {
	options   ListPoliciesPaginatorOptions
	client    ListPoliciesAPIClient
	params    *ListPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListPoliciesPaginator returns a new ListPoliciesPaginator
func NewListPoliciesPaginator(client ListPoliciesAPIClient, params *ListPoliciesInput, optFns ...func(*ListPoliciesPaginatorOptions)) *ListPoliciesPaginator {
	if params == nil {
		params = &ListPoliciesInput{}
	}

	options := ListPoliciesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPolicies page.
func (p *ListPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPoliciesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListPolicies(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListPoliciesAPIClient is a client that implements the ListPolicies operation.
type ListPoliciesAPIClient interface {
	ListPolicies(context.Context, *ListPoliciesInput, ...func(*Options)) (*ListPoliciesOutput, error)
}

var _ ListPoliciesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPolicies",
	}
}
