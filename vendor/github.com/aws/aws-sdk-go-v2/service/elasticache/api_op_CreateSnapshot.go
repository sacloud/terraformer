// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CreateSnapshot operation.
type CreateSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The identifier of an existing cluster. The snapshot is created from this
	// cluster.
	CacheClusterId *string `type:"string"`

	// The ID of the KMS key used to encrypt the snapshot.
	KmsKeyId *string `type:"string"`

	// The identifier of an existing replication group. The snapshot is created
	// from this replication group.
	ReplicationGroupId *string `type:"string"`

	// A name for the snapshot being created.
	//
	// SnapshotName is a required field
	SnapshotName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSnapshotInput"}

	if s.SnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Represents a copy of an entire Redis cluster as of the time when the snapshot
	// was taken.
	Snapshot *Snapshot `type:"structure"`
}

// String returns the string representation
func (s CreateSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSnapshot = "CreateSnapshot"

// CreateSnapshotRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Creates a copy of an entire cluster or replication group at a specific moment
// in time.
//
// This operation is valid for Redis only.
//
//    // Example sending a request using CreateSnapshotRequest.
//    req := client.CreateSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CreateSnapshot
func (c *Client) CreateSnapshotRequest(input *CreateSnapshotInput) CreateSnapshotRequest {
	op := &aws.Operation{
		Name:       opCreateSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSnapshotInput{}
	}

	req := c.newRequest(op, input, &CreateSnapshotOutput{})
	return CreateSnapshotRequest{Request: req, Input: input, Copy: c.CreateSnapshotRequest}
}

// CreateSnapshotRequest is the request type for the
// CreateSnapshot API operation.
type CreateSnapshotRequest struct {
	*aws.Request
	Input *CreateSnapshotInput
	Copy  func(*CreateSnapshotInput) CreateSnapshotRequest
}

// Send marshals and sends the CreateSnapshot API request.
func (r CreateSnapshotRequest) Send(ctx context.Context) (*CreateSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSnapshotResponse{
		CreateSnapshotOutput: r.Request.Data.(*CreateSnapshotOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSnapshotResponse is the response type for the
// CreateSnapshot API operation.
type CreateSnapshotResponse struct {
	*CreateSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSnapshot request.
func (r *CreateSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
