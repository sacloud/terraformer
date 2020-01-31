// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

type PutBucketReplicationInput struct {
	_ struct{} `type:"structure" payload:"ReplicationConfiguration"`

	// The name of the bucket
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// A container for replication rules. You can add up to 1,000 rules. The maximum
	// size of a replication configuration is 2 MB.
	//
	// ReplicationConfiguration is a required field
	ReplicationConfiguration *ReplicationConfiguration `locationName:"ReplicationConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	Token *string `location:"header" locationName:"x-amz-bucket-object-lock-token" type:"string"`
}

// String returns the string representation
func (s PutBucketReplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketReplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketReplicationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.ReplicationConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationConfiguration"))
	}
	if s.ReplicationConfiguration != nil {
		if err := s.ReplicationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ReplicationConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketReplicationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketReplicationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Token != nil {
		v := *s.Token

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-bucket-object-lock-token", protocol.StringValue(v), metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.ReplicationConfiguration != nil {
		v := s.ReplicationConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "ReplicationConfiguration", v, metadata)
	}
	return nil
}

type PutBucketReplicationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketReplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketReplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketReplication = "PutBucketReplication"

// PutBucketReplicationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Creates a replication configuration or replaces an existing one. For more
// information, see Replication (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html)
// in the Amazon S3 Developer Guide.
//
// To perform this operation, the user or role performing the operation must
// have the iam:PassRole (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html)
// permission.
//
// Specify the replication configuration in the request body. In the replication
// configuration, you provide the name of the destination bucket where you want
// Amazon S3 to replicate objects, the IAM role that Amazon S3 can assume to
// replicate objects on your behalf, and other relevant information.
//
// A replication configuration must include at least one rule, and can contain
// a maximum of 1,000. Each rule identifies a subset of objects to replicate
// by filtering the objects in the source bucket. To choose additional subsets
// of objects to replicate, add a rule for each subset. All rules must specify
// the same destination bucket.
//
// To specify a subset of the objects in the source bucket to apply a replication
// rule to, add the Filter element as a child of the Rule element. You can filter
// objects based on an object key prefix, one or more object tags, or both.
// When you add the Filter element in the configuration, you must also add the
// following elements: DeleteMarkerReplication, Status, and Priority.
//
// For information about enabling versioning on a bucket, see Using Versioning
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html).
//
// By default, a resource owner, in this case the AWS account that created the
// bucket, can perform this operation. The resource owner can also grant others
// permissions to perform the operation. For more information about permissions,
// see Specifying Permissions in a Policy (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
//
// Handling Replication of Encrypted Objects
//
// By default, Amazon S3 doesn't replicate objects that are stored at rest using
// server-side encryption with CMKs stored in AWS KMS. To replicate AWS KMS-encrypted
// objects, add the following: SourceSelectionCriteria, SseKmsEncryptedObjects,
// Status, EncryptionConfiguration, and ReplicaKmsKeyID. For information about
// replication configuration, see Replicating Objects Created with SSE Using
// CMKs stored in AWS KMS (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-config-for-kms-objects.html).
//
// PutBucketReplication has the following special errors:
//
//    * Error code: InvalidRequest Description: If the <Owner> in <AccessControlTranslation>
//    has a value, the <Account> element must be specified. HTTP 400
//
//    * Error code: InvalidArgument Description: The <Account> element is empty.
//    It must contain a valid account ID. HTTP 400
//
//    * Error code: InvalidArgument Description: The AWS account specified in
//    the <Account> element must match the destination bucket owner. HTTP 400
//
// The following operations are related to PutBucketReplication:
//
//    * GetBucketReplication
//
//    * DeleteBucketReplication
//
//    // Example sending a request using PutBucketReplicationRequest.
//    req := client.PutBucketReplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketReplication
func (c *Client) PutBucketReplicationRequest(input *PutBucketReplicationInput) PutBucketReplicationRequest {
	op := &aws.Operation{
		Name:       opPutBucketReplication,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?replication",
	}

	if input == nil {
		input = &PutBucketReplicationInput{}
	}

	req := c.newRequest(op, input, &PutBucketReplicationOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketReplicationRequest{Request: req, Input: input, Copy: c.PutBucketReplicationRequest}
}

// PutBucketReplicationRequest is the request type for the
// PutBucketReplication API operation.
type PutBucketReplicationRequest struct {
	*aws.Request
	Input *PutBucketReplicationInput
	Copy  func(*PutBucketReplicationInput) PutBucketReplicationRequest
}

// Send marshals and sends the PutBucketReplication API request.
func (r PutBucketReplicationRequest) Send(ctx context.Context) (*PutBucketReplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketReplicationResponse{
		PutBucketReplicationOutput: r.Request.Data.(*PutBucketReplicationOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketReplicationResponse is the response type for the
// PutBucketReplication API operation.
type PutBucketReplicationResponse struct {
	*PutBucketReplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketReplication request.
func (r *PutBucketReplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
