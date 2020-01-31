// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RemoveListenerCertificatesInput struct {
	_ struct{} `type:"structure"`

	// The certificate to remove. You can specify one certificate per call. Set
	// CertificateArn to the certificate ARN but do not set IsDefault.
	//
	// Certificates is a required field
	Certificates []Certificate `type:"list" required:"true"`

	// The Amazon Resource Name (ARN) of the listener.
	//
	// ListenerArn is a required field
	ListenerArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveListenerCertificatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveListenerCertificatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveListenerCertificatesInput"}

	if s.Certificates == nil {
		invalidParams.Add(aws.NewErrParamRequired("Certificates"))
	}

	if s.ListenerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ListenerArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RemoveListenerCertificatesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveListenerCertificatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemoveListenerCertificates = "RemoveListenerCertificates"

// RemoveListenerCertificatesRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Removes the specified certificate from the certificate list for the specified
// HTTPS or TLS listener.
//
// You can't remove the default certificate for a listener. To replace the default
// certificate, call ModifyListener.
//
// To list the certificates for your listener, use DescribeListenerCertificates.
//
//    // Example sending a request using RemoveListenerCertificatesRequest.
//    req := client.RemoveListenerCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancingv2-2015-12-01/RemoveListenerCertificates
func (c *Client) RemoveListenerCertificatesRequest(input *RemoveListenerCertificatesInput) RemoveListenerCertificatesRequest {
	op := &aws.Operation{
		Name:       opRemoveListenerCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveListenerCertificatesInput{}
	}

	req := c.newRequest(op, input, &RemoveListenerCertificatesOutput{})
	return RemoveListenerCertificatesRequest{Request: req, Input: input, Copy: c.RemoveListenerCertificatesRequest}
}

// RemoveListenerCertificatesRequest is the request type for the
// RemoveListenerCertificates API operation.
type RemoveListenerCertificatesRequest struct {
	*aws.Request
	Input *RemoveListenerCertificatesInput
	Copy  func(*RemoveListenerCertificatesInput) RemoveListenerCertificatesRequest
}

// Send marshals and sends the RemoveListenerCertificates API request.
func (r RemoveListenerCertificatesRequest) Send(ctx context.Context) (*RemoveListenerCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveListenerCertificatesResponse{
		RemoveListenerCertificatesOutput: r.Request.Data.(*RemoveListenerCertificatesOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveListenerCertificatesResponse is the response type for the
// RemoveListenerCertificates API operation.
type RemoveListenerCertificatesResponse struct {
	*RemoveListenerCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveListenerCertificates request.
func (r *RemoveListenerCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
