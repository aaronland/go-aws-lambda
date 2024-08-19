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

// Creates a Lambda function. To create a function, you need a [deployment package] and an [execution role]. The
// deployment package is a .zip file archive or container image that contains your
// function code. The execution role grants the function permission to use Amazon
// Web Services, such as Amazon CloudWatch Logs for log streaming and X-Ray for
// request tracing.
//
// If the deployment package is a [container image], then you set the package type to Image . For a
// container image, the code property must include the URI of a container image in
// the Amazon ECR registry. You do not need to specify the handler and runtime
// properties.
//
// If the deployment package is a [.zip file archive], then you set the package type to Zip . For a
// .zip file archive, the code property specifies the location of the .zip file.
// You must also specify the handler and runtime properties. The code in the
// deployment package must be compatible with the target instruction set
// architecture of the function ( x86-64 or arm64 ). If you do not specify the
// architecture, then the default value is x86-64 .
//
// When you create a function, Lambda provisions an instance of the function and
// its supporting resources. If your function connects to a VPC, this process can
// take a minute or so. During this time, you can't invoke or modify the function.
// The State , StateReason , and StateReasonCode fields in the response from GetFunctionConfiguration
// indicate when the function is ready to invoke. For more information, see [Lambda function states].
//
// A function has an unpublished version, and can have published versions and
// aliases. The unpublished version changes when you update your function's code
// and configuration. A published version is a snapshot of your function code and
// configuration that can't be changed. An alias is a named resource that maps to a
// version, and can be changed to map to a different version. Use the Publish
// parameter to create version 1 of your function from its initial configuration.
//
// The other parameters let you configure version-specific and function-level
// settings. You can modify version-specific settings later with UpdateFunctionConfiguration. Function-level
// settings apply to both the unpublished and published versions of the function,
// and include tags (TagResource ) and per-function concurrency limits (PutFunctionConcurrency ).
//
// You can use code signing if your deployment package is a .zip file archive. To
// enable code signing for this function, specify the ARN of a code-signing
// configuration. When a user attempts to deploy a code package with UpdateFunctionCode, Lambda
// checks that the code package has a valid signature from a trusted publisher. The
// code-signing configuration includes set of signing profiles, which define the
// trusted publishers for this function.
//
// If another Amazon Web Services account or an Amazon Web Service invokes your
// function, use AddPermissionto grant permission by creating a resource-based Identity and
// Access Management (IAM) policy. You can grant permissions at the function level,
// on a version, or on an alias.
//
// To invoke your function directly, use Invoke. To invoke your function in response to
// events in other Amazon Web Services, create an event source mapping (CreateEventSourceMapping ), or
// configure a function trigger in the other service. For more information, see [Invoking Lambda functions].
//
// [Invoking Lambda functions]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-invocation.html
// [Lambda function states]: https://docs.aws.amazon.com/lambda/latest/dg/functions-states.html
// [.zip file archive]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html#gettingstarted-package-zip
// [execution role]: https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html#lambda-intro-execution-role
// [container image]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html
// [deployment package]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html
func (c *Client) CreateFunction(ctx context.Context, params *CreateFunctionInput, optFns ...func(*Options)) (*CreateFunctionOutput, error) {
	if params == nil {
		params = &CreateFunctionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFunction", params, optFns, c.addOperationCreateFunctionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFunctionInput struct {

	// The code for the function.
	//
	// This member is required.
	Code *types.FunctionCode

	// The name or ARN of the Lambda function.
	//
	// Name formats
	//
	//   - Function name – my-function .
	//
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//
	//   - Partial ARN – 123456789012:function:my-function .
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// The Amazon Resource Name (ARN) of the function's execution role.
	//
	// This member is required.
	Role *string

	// The instruction set architecture that the function supports. Enter a string
	// array with one of the valid values (arm64 or x86_64). The default value is
	// x86_64 .
	Architectures []types.Architecture

	// To enable code signing for this function, specify the ARN of a code-signing
	// configuration. A code-signing configuration includes a set of signing profiles,
	// which define the trusted publishers for this function.
	CodeSigningConfigArn *string

	// A dead-letter queue configuration that specifies the queue or topic where
	// Lambda sends asynchronous events when they fail processing. For more
	// information, see [Dead-letter queues].
	//
	// [Dead-letter queues]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-dlq
	DeadLetterConfig *types.DeadLetterConfig

	// A description of the function.
	Description *string

	// Environment variables that are accessible from function code during execution.
	Environment *types.Environment

	// The size of the function's /tmp directory in MB. The default value is 512, but
	// can be any whole number between 512 and 10,240 MB. For more information, see [Configuring ephemeral storage (console)].
	//
	// [Configuring ephemeral storage (console)]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-ephemeral-storage
	EphemeralStorage *types.EphemeralStorage

	// Connection settings for an Amazon EFS file system.
	FileSystemConfigs []types.FileSystemConfig

	// The name of the method within your code that Lambda calls to run your function.
	// Handler is required if the deployment package is a .zip file archive. The format
	// includes the file name. It can also include namespaces and other qualifiers,
	// depending on the runtime. For more information, see [Lambda programming model].
	//
	// [Lambda programming model]: https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html
	Handler *string

	// Container image [configuration values] that override the values in the container image Dockerfile.
	//
	// [configuration values]: https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms
	ImageConfig *types.ImageConfig

	// The ARN of the Key Management Service (KMS) customer managed key that's used to
	// encrypt your function's [environment variables]. When [Lambda SnapStart] is activated, Lambda also uses this key is to
	// encrypt your function's snapshot. If you deploy your function using a container
	// image, Lambda also uses this key to encrypt your function when it's deployed.
	// Note that this is not the same key that's used to protect your container image
	// in the Amazon Elastic Container Registry (Amazon ECR). If you don't provide a
	// customer managed key, Lambda uses a default service key.
	//
	// [Lambda SnapStart]: https://docs.aws.amazon.com/lambda/latest/dg/snapstart-security.html
	// [environment variables]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-encryption
	KMSKeyArn *string

	// A list of [function layers] to add to the function's execution environment. Specify each layer
	// by its ARN, including the version.
	//
	// [function layers]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html
	Layers []string

	// The function's Amazon CloudWatch Logs configuration settings.
	LoggingConfig *types.LoggingConfig

	// The amount of [memory available to the function] at runtime. Increasing the function memory also increases its
	// CPU allocation. The default value is 128 MB. The value can be any multiple of 1
	// MB.
	//
	// [memory available to the function]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console
	MemorySize *int32

	// The type of deployment package. Set to Image for container image and set to Zip
	// for .zip file archive.
	PackageType types.PackageType

	// Set to true to publish the first version of the function during creation.
	Publish bool

	// The identifier of the function's [runtime]. Runtime is required if the deployment
	// package is a .zip file archive. Specifying a runtime results in an error if
	// you're deploying a function using a container image.
	//
	// The following list includes deprecated runtimes. Lambda blocks creating new
	// functions and updating existing functions shortly after each runtime is
	// deprecated. For more information, see [Runtime use after deprecation].
	//
	// For a list of all currently supported runtimes, see [Supported runtimes].
	//
	// [Runtime use after deprecation]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-deprecation-levels
	// [runtime]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html
	// [Supported runtimes]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtimes-supported
	Runtime types.Runtime

	// The function's [SnapStart] setting.
	//
	// [SnapStart]: https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html
	SnapStart *types.SnapStart

	// A list of [tags] to apply to the function.
	//
	// [tags]: https://docs.aws.amazon.com/lambda/latest/dg/tagging.html
	Tags map[string]string

	// The amount of time (in seconds) that Lambda allows a function to run before
	// stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds.
	// For more information, see [Lambda execution environment].
	//
	// [Lambda execution environment]: https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html
	Timeout *int32

	// Set Mode to Active to sample and trace a subset of incoming requests with [X-Ray].
	//
	// [X-Ray]: https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html
	TracingConfig *types.TracingConfig

	// For network connectivity to Amazon Web Services resources in a VPC, specify a
	// list of security groups and subnets in the VPC. When you connect a function to a
	// VPC, it can access resources and the internet only through that VPC. For more
	// information, see [Configuring a Lambda function to access resources in a VPC].
	//
	// [Configuring a Lambda function to access resources in a VPC]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

// Details about a function's configuration.
type CreateFunctionOutput struct {

	// The instruction set architecture that the function supports. Architecture is a
	// string array with one of the valid values. The default architecture value is
	// x86_64 .
	Architectures []types.Architecture

	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string

	// The size of the function's deployment package, in bytes.
	CodeSize int64

	// The function's dead letter queue.
	DeadLetterConfig *types.DeadLetterConfig

	// The function's description.
	Description *string

	// The function's [environment variables]. Omitted from CloudTrail logs.
	//
	// [environment variables]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html
	Environment *types.EnvironmentResponse

	// The size of the function's /tmp directory in MB. The default value is 512, but
	// can be any whole number between 512 and 10,240 MB. For more information, see [Configuring ephemeral storage (console)].
	//
	// [Configuring ephemeral storage (console)]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-ephemeral-storage
	EphemeralStorage *types.EphemeralStorage

	// Connection settings for an [Amazon EFS file system].
	//
	// [Amazon EFS file system]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-filesystem.html
	FileSystemConfigs []types.FileSystemConfig

	// The function's Amazon Resource Name (ARN).
	FunctionArn *string

	// The name of the function.
	FunctionName *string

	// The function that Lambda calls to begin running your function.
	Handler *string

	// The function's image configuration values.
	ImageConfigResponse *types.ImageConfigResponse

	// The KMS key that's used to encrypt the function's [environment variables]. When [Lambda SnapStart] is activated, this
	// key is also used to encrypt the function's snapshot. This key is returned only
	// if you've configured a customer managed key.
	//
	// [Lambda SnapStart]: https://docs.aws.amazon.com/lambda/latest/dg/snapstart-security.html
	// [environment variables]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-encryption
	KMSKeyArn *string

	// The date and time that the function was last updated, in [ISO-8601 format]
	// (YYYY-MM-DDThh:mm:ss.sTZD).
	//
	// [ISO-8601 format]: https://www.w3.org/TR/NOTE-datetime
	LastModified *string

	// The status of the last update that was performed on the function. This is first
	// set to Successful after function creation completes.
	LastUpdateStatus types.LastUpdateStatus

	// The reason for the last update that was performed on the function.
	LastUpdateStatusReason *string

	// The reason code for the last update that was performed on the function.
	LastUpdateStatusReasonCode types.LastUpdateStatusReasonCode

	// The function's [layers].
	//
	// [layers]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html
	Layers []types.Layer

	// The function's Amazon CloudWatch Logs configuration settings.
	LoggingConfig *types.LoggingConfig

	// For Lambda@Edge functions, the ARN of the main function.
	MasterArn *string

	// The amount of memory available to the function at runtime.
	MemorySize *int32

	// The type of deployment package. Set to Image for container image and set Zip
	// for .zip file archive.
	PackageType types.PackageType

	// The latest updated revision of the function or alias.
	RevisionId *string

	// The function's execution role.
	Role *string

	// The identifier of the function's [runtime]. Runtime is required if the deployment
	// package is a .zip file archive. Specifying a runtime results in an error if
	// you're deploying a function using a container image.
	//
	// The following list includes deprecated runtimes. Lambda blocks creating new
	// functions and updating existing functions shortly after each runtime is
	// deprecated. For more information, see [Runtime use after deprecation].
	//
	// For a list of all currently supported runtimes, see [Supported runtimes].
	//
	// [Runtime use after deprecation]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-deprecation-levels
	// [runtime]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html
	// [Supported runtimes]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtimes-supported
	Runtime types.Runtime

	// The ARN of the runtime and any errors that occured.
	RuntimeVersionConfig *types.RuntimeVersionConfig

	// The ARN of the signing job.
	SigningJobArn *string

	// The ARN of the signing profile version.
	SigningProfileVersionArn *string

	// Set ApplyOn to PublishedVersions to create a snapshot of the initialized
	// execution environment when you publish a function version. For more information,
	// see [Improving startup performance with Lambda SnapStart].
	//
	// [Improving startup performance with Lambda SnapStart]: https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html
	SnapStart *types.SnapStartResponse

	// The current state of the function. When the state is Inactive , you can
	// reactivate the function by invoking it.
	State types.State

	// The reason for the function's current state.
	StateReason *string

	// The reason code for the function's current state. When the code is Creating ,
	// you can't invoke or modify the function.
	StateReasonCode types.StateReasonCode

	// The amount of time in seconds that Lambda allows a function to run before
	// stopping it.
	Timeout *int32

	// The function's X-Ray tracing configuration.
	TracingConfig *types.TracingConfigResponse

	// The version of the Lambda function.
	Version *string

	// The function's networking configuration.
	VpcConfig *types.VpcConfigResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFunctionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFunction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFunction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFunction"); err != nil {
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
	if err = addOpCreateFunctionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFunction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFunction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFunction",
	}
}