// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package budgets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request of DescribeBudgets
type DescribeBudgetsInput struct {
	_ struct{} `type:"structure"`

	// The accountId that is associated with the budgets that you want descriptions
	// of.
	//
	// AccountId is a required field
	AccountId *string `min:"12" type:"string" required:"true"`

	// An optional integer that represents how many entries a paginated response
	// contains. The maximum is 100.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token that you include in your request to indicate the next
	// set of results that you want to retrieve.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeBudgetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBudgetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBudgetsInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.AccountId != nil && len(*s.AccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountId", 12))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response of DescribeBudgets
type DescribeBudgetsOutput struct {
	_ struct{} `type:"structure"`

	// A list of budgets.
	Budgets []Budget `type:"list"`

	// The pagination token in the service response that indicates the next set
	// of results that you can retrieve.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeBudgetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeBudgets = "DescribeBudgets"

// DescribeBudgetsRequest returns a request value for making API operation for
// AWS Budgets.
//
// Lists the budgets that are associated with an account.
//
// The Request Syntax section shows the BudgetLimit syntax. For PlannedBudgetLimits,
// see the Examples (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_DescribeBudgets.html#API_DescribeBudgets_Examples)
// section.
//
//    // Example sending a request using DescribeBudgetsRequest.
//    req := client.DescribeBudgetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeBudgetsRequest(input *DescribeBudgetsInput) DescribeBudgetsRequest {
	op := &aws.Operation{
		Name:       opDescribeBudgets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBudgetsInput{}
	}

	req := c.newRequest(op, input, &DescribeBudgetsOutput{})
	return DescribeBudgetsRequest{Request: req, Input: input, Copy: c.DescribeBudgetsRequest}
}

// DescribeBudgetsRequest is the request type for the
// DescribeBudgets API operation.
type DescribeBudgetsRequest struct {
	*aws.Request
	Input *DescribeBudgetsInput
	Copy  func(*DescribeBudgetsInput) DescribeBudgetsRequest
}

// Send marshals and sends the DescribeBudgets API request.
func (r DescribeBudgetsRequest) Send(ctx context.Context) (*DescribeBudgetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeBudgetsResponse{
		DescribeBudgetsOutput: r.Request.Data.(*DescribeBudgetsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeBudgetsResponse is the response type for the
// DescribeBudgets API operation.
type DescribeBudgetsResponse struct {
	*DescribeBudgetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeBudgets request.
func (r *DescribeBudgetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
