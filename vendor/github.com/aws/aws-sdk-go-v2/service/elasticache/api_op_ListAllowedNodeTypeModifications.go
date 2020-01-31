// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input parameters for the ListAllowedNodeTypeModifications operation.
type ListAllowedNodeTypeModificationsInput struct {
	_ struct{} `type:"structure"`

	// The name of the cluster you want to scale up to a larger node instanced type.
	// ElastiCache uses the cluster id to identify the current node type of this
	// cluster and from that to create a list of node types you can scale up to.
	//
	// You must provide a value for either the CacheClusterId or the ReplicationGroupId.
	CacheClusterId *string `type:"string"`

	// The name of the replication group want to scale up to a larger node type.
	// ElastiCache uses the replication group id to identify the current node type
	// being used by this replication group, and from that to create a list of node
	// types you can scale up to.
	//
	// You must provide a value for either the CacheClusterId or the ReplicationGroupId.
	ReplicationGroupId *string `type:"string"`
}

// String returns the string representation
func (s ListAllowedNodeTypeModificationsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the allowed node types you can use to modify your cluster or replication
// group.
type ListAllowedNodeTypeModificationsOutput struct {
	_ struct{} `type:"structure"`

	// A string list, each element of which specifies a cache node type which you
	// can use to scale your cluster or replication group.
	//
	// When scaling down on a Redis cluster or replication group using ModifyCacheCluster
	// or ModifyReplicationGroup, use a value from this list for the CacheNodeType
	// parameter.
	ScaleDownModifications []string `type:"list"`

	// A string list, each element of which specifies a cache node type which you
	// can use to scale your cluster or replication group.
	//
	// When scaling up a Redis cluster or replication group using ModifyCacheCluster
	// or ModifyReplicationGroup, use a value from this list for the CacheNodeType
	// parameter.
	ScaleUpModifications []string `type:"list"`
}

// String returns the string representation
func (s ListAllowedNodeTypeModificationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAllowedNodeTypeModifications = "ListAllowedNodeTypeModifications"

// ListAllowedNodeTypeModificationsRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Lists all available node types that you can scale your Redis cluster's or
// replication group's current node type.
//
// When you use the ModifyCacheCluster or ModifyReplicationGroup operations
// to scale your cluster or replication group, the value of the CacheNodeType
// parameter must be one of the node types returned by this operation.
//
//    // Example sending a request using ListAllowedNodeTypeModificationsRequest.
//    req := client.ListAllowedNodeTypeModificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/ListAllowedNodeTypeModifications
func (c *Client) ListAllowedNodeTypeModificationsRequest(input *ListAllowedNodeTypeModificationsInput) ListAllowedNodeTypeModificationsRequest {
	op := &aws.Operation{
		Name:       opListAllowedNodeTypeModifications,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAllowedNodeTypeModificationsInput{}
	}

	req := c.newRequest(op, input, &ListAllowedNodeTypeModificationsOutput{})
	return ListAllowedNodeTypeModificationsRequest{Request: req, Input: input, Copy: c.ListAllowedNodeTypeModificationsRequest}
}

// ListAllowedNodeTypeModificationsRequest is the request type for the
// ListAllowedNodeTypeModifications API operation.
type ListAllowedNodeTypeModificationsRequest struct {
	*aws.Request
	Input *ListAllowedNodeTypeModificationsInput
	Copy  func(*ListAllowedNodeTypeModificationsInput) ListAllowedNodeTypeModificationsRequest
}

// Send marshals and sends the ListAllowedNodeTypeModifications API request.
func (r ListAllowedNodeTypeModificationsRequest) Send(ctx context.Context) (*ListAllowedNodeTypeModificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAllowedNodeTypeModificationsResponse{
		ListAllowedNodeTypeModificationsOutput: r.Request.Data.(*ListAllowedNodeTypeModificationsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAllowedNodeTypeModificationsResponse is the response type for the
// ListAllowedNodeTypeModifications API operation.
type ListAllowedNodeTypeModificationsResponse struct {
	*ListAllowedNodeTypeModificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAllowedNodeTypeModifications request.
func (r *ListAllowedNodeTypeModificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
