// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SendMessageBatchInput struct {
	_ struct{} `type:"structure"`

	// A list of SendMessageBatchRequestEntry items.
	//
	// Entries is a required field
	Entries []SendMessageBatchRequestEntry `locationNameList:"SendMessageBatchRequestEntry" type:"list" flattened:"true" required:"true"`

	// The URL of the Amazon SQS queue to which batched messages are sent.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SendMessageBatchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendMessageBatchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendMessageBatchInput"}

	if s.Entries == nil {
		invalidParams.Add(aws.NewErrParamRequired("Entries"))
	}

	if s.QueueUrl == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueueUrl"))
	}
	if s.Entries != nil {
		for i, v := range s.Entries {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Entries", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// For each message in the batch, the response contains a SendMessageBatchResultEntry
// tag if the message succeeds or a BatchResultErrorEntry tag if the message
// fails.
type SendMessageBatchOutput struct {
	_ struct{} `type:"structure"`

	// A list of BatchResultErrorEntry items with error details about each message
	// that can't be enqueued.
	//
	// Failed is a required field
	Failed []BatchResultErrorEntry `locationNameList:"BatchResultErrorEntry" type:"list" flattened:"true" required:"true"`

	// A list of SendMessageBatchResultEntry items.
	//
	// Successful is a required field
	Successful []SendMessageBatchResultEntry `locationNameList:"SendMessageBatchResultEntry" type:"list" flattened:"true" required:"true"`
}

// String returns the string representation
func (s SendMessageBatchOutput) String() string {
	return awsutil.Prettify(s)
}

const opSendMessageBatch = "SendMessageBatch"

// SendMessageBatchRequest returns a request value for making API operation for
// Amazon Simple Queue Service.
//
// Delivers up to ten messages to the specified queue. This is a batch version
// of SendMessage. For a FIFO queue, multiple messages within a single batch
// are enqueued in the order they are sent.
//
// The result of sending each message is reported individually in the response.
// Because the batch request can result in a combination of successful and unsuccessful
// actions, you should check for batch errors even when the call returns an
// HTTP status code of 200.
//
// The maximum allowed individual message size and the maximum total payload
// size (the sum of the individual lengths of all of the batched messages) are
// both 256 KB (262,144 bytes).
//
// A message can include only XML, JSON, and unformatted text. The following
// Unicode characters are allowed:
//
// #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF
//
// Any characters not included in this list will be rejected. For more information,
// see the W3C specification for characters (http://www.w3.org/TR/REC-xml/#charsets).
//
// If you don't specify the DelaySeconds parameter for an entry, Amazon SQS
// uses the default value for the queue.
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example,
// a parameter list with two elements looks like this:
//
// &Attribute.1=first
//
// &Attribute.2=second
//
//    // Example sending a request using SendMessageBatchRequest.
//    req := client.SendMessageBatchRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/SendMessageBatch
func (c *Client) SendMessageBatchRequest(input *SendMessageBatchInput) SendMessageBatchRequest {
	op := &aws.Operation{
		Name:       opSendMessageBatch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendMessageBatchInput{}
	}

	req := c.newRequest(op, input, &SendMessageBatchOutput{})
	return SendMessageBatchRequest{Request: req, Input: input, Copy: c.SendMessageBatchRequest}
}

// SendMessageBatchRequest is the request type for the
// SendMessageBatch API operation.
type SendMessageBatchRequest struct {
	*aws.Request
	Input *SendMessageBatchInput
	Copy  func(*SendMessageBatchInput) SendMessageBatchRequest
}

// Send marshals and sends the SendMessageBatch API request.
func (r SendMessageBatchRequest) Send(ctx context.Context) (*SendMessageBatchResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendMessageBatchResponse{
		SendMessageBatchOutput: r.Request.Data.(*SendMessageBatchOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendMessageBatchResponse is the response type for the
// SendMessageBatch API operation.
type SendMessageBatchResponse struct {
	*SendMessageBatchOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendMessageBatch request.
func (r *SendMessageBatchResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
