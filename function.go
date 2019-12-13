package lambda

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/aaronland/go-aws-session"
	"github.com/aws/aws-sdk-go/aws"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	aws_lambda "github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaFunction struct {
	service   *aws_lambda.Lambda
	func_name string
	func_type string
}

func NewLambdaFunctionWithDSN(dsn string, func_name string, func_type string) (*LambdaFunction, error) {

	sess, err := session.NewSessionWithDSN(dsn)

	if err != nil {
		return nil, err
	}

	return NewLambdaFunctionWithSession(sess, func_name, func_type)
}

func NewLambdaFunctionWithSession(sess *aws_session.Session, func_name string, func_type string) (*LambdaFunction, error) {

	svc := aws_lambda.New(sess)
	return NewLambdaFunctionWithService(svc, func_name, func_type)
}

func NewLambdaFunctionWithService(svc *aws_lambda.Lambda, func_name string, func_type string) (*LambdaFunction, error) {

	f := &LambdaFunction{
		service:   svc,
		func_name: func_name,
		func_type: func_type,
	}

	return f, nil
}

func (f *LambdaFunction) Invoke(ctx context.Context, payload interface{}) (*aws_lambda.InvokeOutput, error) {

	enc_payload, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	input := &aws_lambda.InvokeInput{
		FunctionName:   aws.String(f.func_name),
		InvocationType: aws.String(f.func_type),
		Payload:        enc_payload,
	}

	if *input.InvocationType == "RequestResponse" {
		input.LogType = aws.String("Tail")
	}

	rsp, err := f.service.Invoke(input)

	if err != nil {
		return nil, err
	}

	if *input.InvocationType != "RequestResponse" {
		return nil, nil
	}

	enc_result := *rsp.LogResult

	result, err := base64.StdEncoding.DecodeString(enc_result)

	if err != nil {
		return nil, err
	}

	if *rsp.StatusCode != 200 {
		return nil, errors.New(string(result))
	}

	return rsp, nil
}
