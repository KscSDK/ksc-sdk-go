// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package billunion

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opDescribeBillSummaryByPayMode = "DescribeBillSummaryByPayMode"

// DescribeBillSummaryByPayModeRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeBillSummaryByPayMode operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeBillSummaryByPayMode for more information on using the DescribeBillSummaryByPayMode
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeBillSummaryByPayModeRequest method.
//    req, resp := client.DescribeBillSummaryByPayModeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/bill-union-2020-01-01/DescribeBillSummaryByPayMode
func (c *BillUnion) DescribeBillSummaryByPayModeRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBillSummaryByPayMode,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeBillSummaryByPayMode API operation for bill-union.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for bill-union's
// API operation DescribeBillSummaryByPayMode for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/bill-union-2020-01-01/DescribeBillSummaryByPayMode
func (c *BillUnion) DescribeBillSummaryByPayMode(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBillSummaryByPayModeRequest(input)
	return out, req.Send()
}

// DescribeBillSummaryByPayModeWithContext is the same as DescribeBillSummaryByPayMode with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBillSummaryByPayMode for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BillUnion) DescribeBillSummaryByPayModeWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBillSummaryByPayModeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBillSummaryByProduct = "DescribeBillSummaryByProduct"

// DescribeBillSummaryByProductRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeBillSummaryByProduct operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeBillSummaryByProduct for more information on using the DescribeBillSummaryByProduct
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeBillSummaryByProductRequest method.
//    req, resp := client.DescribeBillSummaryByProductRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/bill-union-2020-01-01/DescribeBillSummaryByProduct
func (c *BillUnion) DescribeBillSummaryByProductRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBillSummaryByProduct,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeBillSummaryByProduct API operation for bill-union.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for bill-union's
// API operation DescribeBillSummaryByProduct for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/bill-union-2020-01-01/DescribeBillSummaryByProduct
func (c *BillUnion) DescribeBillSummaryByProduct(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBillSummaryByProductRequest(input)
	return out, req.Send()
}

// DescribeBillSummaryByProductWithContext is the same as DescribeBillSummaryByProduct with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBillSummaryByProduct for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BillUnion) DescribeBillSummaryByProductWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBillSummaryByProductRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBillSummaryByProject = "DescribeBillSummaryByProject"

// DescribeBillSummaryByProjectRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeBillSummaryByProject operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeBillSummaryByProject for more information on using the DescribeBillSummaryByProject
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeBillSummaryByProjectRequest method.
//    req, resp := client.DescribeBillSummaryByProjectRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/bill-union-2020-01-01/DescribeBillSummaryByProject
func (c *BillUnion) DescribeBillSummaryByProjectRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBillSummaryByProject,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeBillSummaryByProject API operation for bill-union.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for bill-union's
// API operation DescribeBillSummaryByProject for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/bill-union-2020-01-01/DescribeBillSummaryByProject
func (c *BillUnion) DescribeBillSummaryByProject(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBillSummaryByProjectRequest(input)
	return out, req.Send()
}

// DescribeBillSummaryByProjectWithContext is the same as DescribeBillSummaryByProject with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBillSummaryByProject for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BillUnion) DescribeBillSummaryByProjectWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBillSummaryByProjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeInstanceSummaryBills = "DescribeInstanceSummaryBills"

// DescribeInstanceSummaryBillsRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeInstanceSummaryBills operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeInstanceSummaryBills for more information on using the DescribeInstanceSummaryBills
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeInstanceSummaryBillsRequest method.
//    req, resp := client.DescribeInstanceSummaryBillsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/bill-union-2020-01-01/DescribeInstanceSummaryBills
func (c *BillUnion) DescribeInstanceSummaryBillsRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeInstanceSummaryBills,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeInstanceSummaryBills API operation for bill-union.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for bill-union's
// API operation DescribeInstanceSummaryBills for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/bill-union-2020-01-01/DescribeInstanceSummaryBills
func (c *BillUnion) DescribeInstanceSummaryBills(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeInstanceSummaryBillsRequest(input)
	return out, req.Send()
}

// DescribeInstanceSummaryBillsWithContext is the same as DescribeInstanceSummaryBills with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstanceSummaryBills for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BillUnion) DescribeInstanceSummaryBillsWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeInstanceSummaryBillsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeProductCode = "DescribeProductCode"

// DescribeProductCodeRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeProductCode operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeProductCode for more information on using the DescribeProductCode
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeProductCodeRequest method.
//    req, resp := client.DescribeProductCodeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/bill-union-2020-01-01/DescribeProductCode
func (c *BillUnion) DescribeProductCodeRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeProductCode,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeProductCode API operation for bill-union.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for bill-union's
// API operation DescribeProductCode for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/bill-union-2020-01-01/DescribeProductCode
func (c *BillUnion) DescribeProductCode(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeProductCodeRequest(input)
	return out, req.Send()
}

// DescribeProductCodeWithContext is the same as DescribeProductCode with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeProductCode for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BillUnion) DescribeProductCodeWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeProductCodeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}