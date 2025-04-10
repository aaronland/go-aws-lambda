// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The specified code signing configuration does not exist.
type CodeSigningConfigNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *CodeSigningConfigNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CodeSigningConfigNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CodeSigningConfigNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CodeSigningConfigNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *CodeSigningConfigNotFoundException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Your Amazon Web Services account has exceeded its maximum total code size. For
// more information, see [Lambda quotas].
//
// [Lambda quotas]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html
type CodeStorageExceededException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *CodeStorageExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CodeStorageExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CodeStorageExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CodeStorageExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *CodeStorageExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The code signature failed one or more of the validation checks for signature
// mismatch or expiry, and the code signing policy is set to ENFORCE. Lambda blocks
// the deployment.
type CodeVerificationFailedException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *CodeVerificationFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CodeVerificationFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CodeVerificationFailedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CodeVerificationFailedException"
	}
	return *e.ErrorCodeOverride
}
func (e *CodeVerificationFailedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Need additional permissions to configure VPC settings.
type EC2AccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *EC2AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EC2AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EC2AccessDeniedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EC2AccessDeniedException"
	}
	return *e.ErrorCodeOverride
}
func (e *EC2AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Amazon EC2 throttled Lambda during Lambda function initialization using the
// execution role provided for the function.
type EC2ThrottledException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *EC2ThrottledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EC2ThrottledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EC2ThrottledException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EC2ThrottledException"
	}
	return *e.ErrorCodeOverride
}
func (e *EC2ThrottledException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Lambda received an unexpected Amazon EC2 client exception while setting up for
// the Lambda function.
type EC2UnexpectedException struct {
	Message *string

	ErrorCodeOverride *string

	Type         *string
	EC2ErrorCode *string

	noSmithyDocumentSerde
}

func (e *EC2UnexpectedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EC2UnexpectedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EC2UnexpectedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EC2UnexpectedException"
	}
	return *e.ErrorCodeOverride
}
func (e *EC2UnexpectedException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// An error occurred when reading from or writing to a connected file system.
type EFSIOException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *EFSIOException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EFSIOException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EFSIOException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EFSIOException"
	}
	return *e.ErrorCodeOverride
}
func (e *EFSIOException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Lambda function couldn't make a network connection to the configured file
// system.
type EFSMountConnectivityException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *EFSMountConnectivityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EFSMountConnectivityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EFSMountConnectivityException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EFSMountConnectivityException"
	}
	return *e.ErrorCodeOverride
}
func (e *EFSMountConnectivityException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Lambda function couldn't mount the configured file system due to a
// permission or configuration issue.
type EFSMountFailureException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *EFSMountFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EFSMountFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EFSMountFailureException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EFSMountFailureException"
	}
	return *e.ErrorCodeOverride
}
func (e *EFSMountFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Lambda function made a network connection to the configured file system,
// but the mount operation timed out.
type EFSMountTimeoutException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *EFSMountTimeoutException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EFSMountTimeoutException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EFSMountTimeoutException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EFSMountTimeoutException"
	}
	return *e.ErrorCodeOverride
}
func (e *EFSMountTimeoutException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Lambda couldn't create an elastic network interface in the VPC, specified as
// part of Lambda function configuration, because the limit for network interfaces
// has been reached. For more information, see [Lambda quotas].
//
// [Lambda quotas]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html
type ENILimitReachedException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *ENILimitReachedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ENILimitReachedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ENILimitReachedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ENILimitReachedException"
	}
	return *e.ErrorCodeOverride
}
func (e *ENILimitReachedException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The code signature failed the integrity check. If the integrity check fails,
// then Lambda blocks deployment, even if the code signing policy is set to WARN.
type InvalidCodeSignatureException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *InvalidCodeSignatureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCodeSignatureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCodeSignatureException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidCodeSignatureException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidCodeSignatureException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One of the parameters in the request is not valid.
type InvalidParameterValueException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterValueException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request body could not be parsed as JSON, or a request header is invalid.
// For example, the 'x-amzn-RequestId' header is not a valid UUID string.
type InvalidRequestContentException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *InvalidRequestContentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestContentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestContentException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidRequestContentException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRequestContentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The runtime or runtime version specified is not supported.
type InvalidRuntimeException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *InvalidRuntimeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRuntimeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRuntimeException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidRuntimeException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRuntimeException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The security group ID provided in the Lambda function VPC configuration is not
// valid.
type InvalidSecurityGroupIDException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *InvalidSecurityGroupIDException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSecurityGroupIDException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSecurityGroupIDException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSecurityGroupIDException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSecurityGroupIDException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The subnet ID provided in the Lambda function VPC configuration is not valid.
type InvalidSubnetIDException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *InvalidSubnetIDException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSubnetIDException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSubnetIDException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSubnetIDException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSubnetIDException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Lambda could not unzip the deployment package.
type InvalidZipFileException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *InvalidZipFileException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidZipFileException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidZipFileException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidZipFileException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidZipFileException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Lambda couldn't decrypt the environment variables because KMS access was
// denied. Check the Lambda function's KMS permissions.
type KMSAccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *KMSAccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSAccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSAccessDeniedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KMSAccessDeniedException"
	}
	return *e.ErrorCodeOverride
}
func (e *KMSAccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Lambda couldn't decrypt the environment variables because the KMS key used is
// disabled. Check the Lambda function's KMS key settings.
type KMSDisabledException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *KMSDisabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSDisabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSDisabledException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KMSDisabledException"
	}
	return *e.ErrorCodeOverride
}
func (e *KMSDisabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Lambda couldn't decrypt the environment variables because the state of the KMS
// key used is not valid for Decrypt. Check the function's KMS key settings.
type KMSInvalidStateException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *KMSInvalidStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSInvalidStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSInvalidStateException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KMSInvalidStateException"
	}
	return *e.ErrorCodeOverride
}
func (e *KMSInvalidStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Lambda couldn't decrypt the environment variables because the KMS key was not
// found. Check the function's KMS key settings.
type KMSNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *KMSNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KMSNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *KMSNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The permissions policy for the resource is too large. For more information, see [Lambda quotas]
// .
//
// [Lambda quotas]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html
type PolicyLengthExceededException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *PolicyLengthExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PolicyLengthExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PolicyLengthExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PolicyLengthExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *PolicyLengthExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The RevisionId provided does not match the latest RevisionId for the Lambda
// function or alias.
//
//   - For AddPermission and RemovePermission API operations: Call GetPolicy to
//     retrieve the latest RevisionId for your resource.
//
//   - For all other API operations: Call GetFunction or GetAlias to retrieve the
//     latest RevisionId for your resource.
type PreconditionFailedException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *PreconditionFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PreconditionFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PreconditionFailedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PreconditionFailedException"
	}
	return *e.ErrorCodeOverride
}
func (e *PreconditionFailedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified configuration does not exist.
type ProvisionedConcurrencyConfigNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *ProvisionedConcurrencyConfigNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ProvisionedConcurrencyConfigNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ProvisionedConcurrencyConfigNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ProvisionedConcurrencyConfigNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ProvisionedConcurrencyConfigNotFoundException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Lambda has detected your function being invoked in a recursive loop with other
// Amazon Web Services resources and stopped your function's invocation.
type RecursiveInvocationException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *RecursiveInvocationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RecursiveInvocationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RecursiveInvocationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RecursiveInvocationException"
	}
	return *e.ErrorCodeOverride
}
func (e *RecursiveInvocationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request payload exceeded the Invoke request body JSON input quota. For more
// information, see [Lambda quotas].
//
// [Lambda quotas]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html
type RequestTooLargeException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *RequestTooLargeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestTooLargeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestTooLargeException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RequestTooLargeException"
	}
	return *e.ErrorCodeOverride
}
func (e *RequestTooLargeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource already exists, or another operation is in progress.
type ResourceConflictException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *ResourceConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceConflictException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation conflicts with the resource's availability. For example, you
// tried to update an event source mapping in the CREATING state, or you tried to
// delete an event source mapping currently UPDATING.
type ResourceInUseException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceInUseException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource specified in the request does not exist.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The function is inactive and its VPC connection is no longer available. Wait
// for the VPC connection to reestablish and try again.
type ResourceNotReadyException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *ResourceNotReadyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotReadyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotReadyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotReadyException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotReadyException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The Lambda service encountered an internal error.
type ServiceException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *ServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The afterRestore()[runtime hook] encountered an error. For more information, check the Amazon
// CloudWatch logs.
//
// [runtime hook]: https://docs.aws.amazon.com/lambda/latest/dg/snapstart-runtime-hooks.html
type SnapStartException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *SnapStartException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapStartException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapStartException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SnapStartException"
	}
	return *e.ErrorCodeOverride
}
func (e *SnapStartException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Lambda is initializing your function. You can invoke the function when the [function state]
// becomes Active .
//
// [function state]: https://docs.aws.amazon.com/lambda/latest/dg/functions-states.html
type SnapStartNotReadyException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *SnapStartNotReadyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapStartNotReadyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapStartNotReadyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SnapStartNotReadyException"
	}
	return *e.ErrorCodeOverride
}
func (e *SnapStartNotReadyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Lambda couldn't restore the snapshot within the timeout limit.
type SnapStartTimeoutException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *SnapStartTimeoutException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapStartTimeoutException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapStartTimeoutException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SnapStartTimeoutException"
	}
	return *e.ErrorCodeOverride
}
func (e *SnapStartTimeoutException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Lambda couldn't set up VPC access for the Lambda function because one or more
// configured subnets has no available IP addresses.
type SubnetIPAddressLimitReachedException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *SubnetIPAddressLimitReachedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetIPAddressLimitReachedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetIPAddressLimitReachedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetIPAddressLimitReachedException"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetIPAddressLimitReachedException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}

// The request throughput limit was exceeded. For more information, see [Lambda quotas].
//
// [Lambda quotas]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html#api-requests
type TooManyRequestsException struct {
	Message *string

	ErrorCodeOverride *string

	RetryAfterSeconds *string
	Type              *string
	Reason            ThrottleReason

	noSmithyDocumentSerde
}

func (e *TooManyRequestsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRequestsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRequestsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyRequestsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyRequestsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The content type of the Invoke request body is not JSON.
type UnsupportedMediaTypeException struct {
	Message *string

	ErrorCodeOverride *string

	Type *string

	noSmithyDocumentSerde
}

func (e *UnsupportedMediaTypeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedMediaTypeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedMediaTypeException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedMediaTypeException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedMediaTypeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
