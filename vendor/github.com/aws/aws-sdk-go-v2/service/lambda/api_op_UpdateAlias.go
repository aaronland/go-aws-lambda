// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration of a Lambda function [alias].
//
// [alias]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html
func (c *Client) UpdateAlias(ctx context.Context, params *UpdateAliasInput, optFns ...func(*Options)) (*UpdateAliasOutput, error) {
	if params == nil {
		params = &UpdateAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAlias", params, optFns, c.addOperationUpdateAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAliasInput struct {

	// The name or ARN of the Lambda function.
	//
	// Name formats
	//
	//   - Function name - MyFunction .
	//
	//   - Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction .
	//
	//   - Partial ARN - 123456789012:function:MyFunction .
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// The name of the alias.
	//
	// This member is required.
	Name *string

	// A description of the alias.
	Description *string

	// The function version that the alias invokes.
	FunctionVersion *string

	// Only update the alias if the revision ID matches the ID that's specified. Use
	// this option to avoid modifying an alias that has changed since you last read it.
	RevisionId *string

	// The [routing configuration] of the alias.
	//
	// [routing configuration]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html#configuring-alias-routing
	RoutingConfig *types.AliasRoutingConfiguration

	noSmithyDocumentSerde
}

// Provides configuration information about a Lambda function [alias].
//
// [alias]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html
type UpdateAliasOutput struct {

	// The Amazon Resource Name (ARN) of the alias.
	AliasArn *string

	// A description of the alias.
	Description *string

	// The function version that the alias invokes.
	FunctionVersion *string

	// The name of the alias.
	Name *string

	// A unique identifier that changes when you update the alias.
	RevisionId *string

	// The [routing configuration] of the alias.
	//
	// [routing configuration]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html
	RoutingConfig *types.AliasRoutingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAlias{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAlias"); err != nil {
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
	if err = addOpUpdateAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAlias",
	}
}
