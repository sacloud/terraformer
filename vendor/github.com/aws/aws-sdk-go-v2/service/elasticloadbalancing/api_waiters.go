// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilAnyInstanceInService uses the Elastic Load Balancing API operation
// DescribeInstanceHealth to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilAnyInstanceInService(ctx context.Context, input *DescribeInstanceHealthInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilAnyInstanceInService",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "InstanceStates[].State",
				Expected: "InService",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeInstanceHealthInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeInstanceHealthRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilInstanceDeregistered uses the Elastic Load Balancing API operation
// DescribeInstanceHealth to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilInstanceDeregistered(ctx context.Context, input *DescribeInstanceHealthInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilInstanceDeregistered",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "InstanceStates[].State",
				Expected: "OutOfService",
			},
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "InvalidInstance",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeInstanceHealthInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeInstanceHealthRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilInstanceInService uses the Elastic Load Balancing API operation
// DescribeInstanceHealth to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilInstanceInService(ctx context.Context, input *DescribeInstanceHealthInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilInstanceInService",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "InstanceStates[].State",
				Expected: "InService",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "InvalidInstance",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeInstanceHealthInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeInstanceHealthRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
