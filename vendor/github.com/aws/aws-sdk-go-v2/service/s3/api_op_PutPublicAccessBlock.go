// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

type PutPublicAccessBlockInput struct {
	_ struct{} `type:"structure" payload:"PublicAccessBlockConfiguration"`

	// The name of the Amazon S3 bucket whose PublicAccessBlock configuration you
	// want to set.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The PublicAccessBlock configuration that you want to apply to this Amazon
	// S3 bucket. You can enable the configuration options in any combination. For
	// more information about when Amazon S3 considers a bucket or object public,
	// see The Meaning of "Public" (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status)
	// in the Amazon Simple Storage Service Developer Guide.
	//
	// PublicAccessBlockConfiguration is a required field
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `locationName:"PublicAccessBlockConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutPublicAccessBlockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutPublicAccessBlockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutPublicAccessBlockInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.PublicAccessBlockConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("PublicAccessBlockConfiguration"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutPublicAccessBlockInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutPublicAccessBlockInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.PublicAccessBlockConfiguration != nil {
		v := s.PublicAccessBlockConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "PublicAccessBlockConfiguration", v, metadata)
	}
	return nil
}

type PutPublicAccessBlockOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutPublicAccessBlockOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutPublicAccessBlockOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutPublicAccessBlock = "PutPublicAccessBlock"

// PutPublicAccessBlockRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Creates or modifies the PublicAccessBlock configuration for an Amazon S3
// bucket. In order to use this operation, you must have the s3:PutBucketPublicAccessBlock
// permission. For more information about Amazon S3 permissions, see Specifying
// Permissions in a Policy (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html).
//
// When Amazon S3 evaluates the PublicAccessBlock configuration for a bucket
// or an object, it checks the PublicAccessBlock configuration for both the
// bucket (or the bucket that contains the object) and the bucket owner's account.
// If the PublicAccessBlock configurations are different between the bucket
// and the account, Amazon S3 uses the most restrictive combination of the bucket-level
// and account-level settings.
//
// For more information about when Amazon S3 considers a bucket or an object
// public, see The Meaning of "Public" (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status).
//
// Related Resources
//
//    * GetPublicAccessBlock
//
//    * DeletePublicAccessBlock
//
//    * GetBucketPolicyStatus
//
//    * Using Amazon S3 Block Public Access (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html)
//
//    // Example sending a request using PutPublicAccessBlockRequest.
//    req := client.PutPublicAccessBlockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutPublicAccessBlock
func (c *Client) PutPublicAccessBlockRequest(input *PutPublicAccessBlockInput) PutPublicAccessBlockRequest {
	op := &aws.Operation{
		Name:       opPutPublicAccessBlock,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?publicAccessBlock",
	}

	if input == nil {
		input = &PutPublicAccessBlockInput{}
	}

	req := c.newRequest(op, input, &PutPublicAccessBlockOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutPublicAccessBlockRequest{Request: req, Input: input, Copy: c.PutPublicAccessBlockRequest}
}

// PutPublicAccessBlockRequest is the request type for the
// PutPublicAccessBlock API operation.
type PutPublicAccessBlockRequest struct {
	*aws.Request
	Input *PutPublicAccessBlockInput
	Copy  func(*PutPublicAccessBlockInput) PutPublicAccessBlockRequest
}

// Send marshals and sends the PutPublicAccessBlock API request.
func (r PutPublicAccessBlockRequest) Send(ctx context.Context) (*PutPublicAccessBlockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutPublicAccessBlockResponse{
		PutPublicAccessBlockOutput: r.Request.Data.(*PutPublicAccessBlockOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutPublicAccessBlockResponse is the response type for the
// PutPublicAccessBlock API operation.
type PutPublicAccessBlockResponse struct {
	*PutPublicAccessBlockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutPublicAccessBlock request.
func (r *PutPublicAccessBlockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
