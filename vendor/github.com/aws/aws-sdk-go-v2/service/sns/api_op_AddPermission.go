// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type AddPermissionInput struct {
	_ struct{} `type:"structure"`

	// The AWS account IDs of the users (principals) who will be given access to
	// the specified actions. The users must have AWS accounts, but do not need
	// to be signed up for this service.
	//
	// AWSAccountId is a required field
	AWSAccountId []string `type:"list" required:"true"`

	// The action you want to allow for the specified principal(s).
	//
	// Valid values: any Amazon SNS action name.
	//
	// ActionName is a required field
	ActionName []string `type:"list" required:"true"`

	// A unique identifier for the new policy statement.
	//
	// Label is a required field
	Label *string `type:"string" required:"true"`

	// The ARN of the topic whose access control policy you wish to modify.
	//
	// TopicArn is a required field
	TopicArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AddPermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddPermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddPermissionInput"}

	if s.AWSAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AWSAccountId"))
	}

	if s.ActionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActionName"))
	}

	if s.Label == nil {
		invalidParams.Add(aws.NewErrParamRequired("Label"))
	}

	if s.TopicArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TopicArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddPermissionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddPermissionOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddPermission = "AddPermission"

// AddPermissionRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Adds a statement to a topic's access control policy, granting access for
// the specified AWS accounts to the specified actions.
//
//    // Example sending a request using AddPermissionRequest.
//    req := client.AddPermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/AddPermission
func (c *Client) AddPermissionRequest(input *AddPermissionInput) AddPermissionRequest {
	op := &aws.Operation{
		Name:       opAddPermission,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddPermissionInput{}
	}

	req := c.newRequest(op, input, &AddPermissionOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AddPermissionRequest{Request: req, Input: input, Copy: c.AddPermissionRequest}
}

// AddPermissionRequest is the request type for the
// AddPermission API operation.
type AddPermissionRequest struct {
	*aws.Request
	Input *AddPermissionInput
	Copy  func(*AddPermissionInput) AddPermissionRequest
}

// Send marshals and sends the AddPermission API request.
func (r AddPermissionRequest) Send(ctx context.Context) (*AddPermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddPermissionResponse{
		AddPermissionOutput: r.Request.Data.(*AddPermissionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddPermissionResponse is the response type for the
// AddPermission API operation.
type AddPermissionResponse struct {
	*AddPermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddPermission request.
func (r *AddPermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}