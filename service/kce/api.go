// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kce

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opDescribeCluster = "DescribeCluster"

// DescribeClusterRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeCluster operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeCluster for more information on using the DescribeCluster
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeClusterRequest method.
//    req, resp := client.DescribeClusterRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/kce-2016-03-04/DescribeCluster
func (c *Kce) DescribeClusterRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCluster,
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

// DescribeCluster API operation for kce.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for kce's
// API operation DescribeCluster for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/kce-2016-03-04/DescribeCluster
func (c *Kce) DescribeCluster(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeClusterRequest(input)
	return out, req.Send()
}

// DescribeClusterWithContext is the same as DescribeCluster with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCluster for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Kce) DescribeClusterWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeClusterRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}