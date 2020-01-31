// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListDomainNamesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ListDomainNamesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDomainNamesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	return nil
}

// The result of a ListDomainNames operation. Contains the names of all Elasticsearch
// domains owned by this account.
type ListDomainNamesOutput struct {
	_ struct{} `type:"structure"`

	// List of Elasticsearch domain names.
	DomainNames []DomainInfo `type:"list"`
}

// String returns the string representation
func (s ListDomainNamesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDomainNamesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DomainNames != nil {
		v := s.DomainNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "DomainNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListDomainNames = "ListDomainNames"

// ListDomainNamesRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Returns the name of all Elasticsearch domains owned by the current user's
// account.
//
//    // Example sending a request using ListDomainNamesRequest.
//    req := client.ListDomainNamesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListDomainNamesRequest(input *ListDomainNamesInput) ListDomainNamesRequest {
	op := &aws.Operation{
		Name:       opListDomainNames,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-01-01/domain",
	}

	if input == nil {
		input = &ListDomainNamesInput{}
	}

	req := c.newRequest(op, input, &ListDomainNamesOutput{})
	return ListDomainNamesRequest{Request: req, Input: input, Copy: c.ListDomainNamesRequest}
}

// ListDomainNamesRequest is the request type for the
// ListDomainNames API operation.
type ListDomainNamesRequest struct {
	*aws.Request
	Input *ListDomainNamesInput
	Copy  func(*ListDomainNamesInput) ListDomainNamesRequest
}

// Send marshals and sends the ListDomainNames API request.
func (r ListDomainNamesRequest) Send(ctx context.Context) (*ListDomainNamesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDomainNamesResponse{
		ListDomainNamesOutput: r.Request.Data.(*ListDomainNamesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDomainNamesResponse is the response type for the
// ListDomainNames API operation.
type ListDomainNamesResponse struct {
	*ListDomainNamesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDomainNames request.
func (r *ListDomainNamesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
