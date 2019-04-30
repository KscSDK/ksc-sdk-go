package eip

import (
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

type DescribeAddressesInput struct {
	_ struct{} `type:"structure"`
	AllocationIds []*string `locationName:"AllocationId" flattened:"true"  type:"list"`
}
func (s *DescribeAddressesInput) String() string{
	return awsutil.Prettify(s)
}

func (s *DescribeAddressesInput) GoString() string{
	return s.String()
}
func (s *DescribeAddressesInput) SetAllocationIds(v []*string) *DescribeAddressesInput {
	s.AllocationIds = v
	return s
}

type Address struct {
	_ struct{} `type:"structure"`
	CreateTime *interface{} `json:"CreateTime"`
	ProjectId *interface{} `json:"ProjectId"`
	PublicIp *interface{} `json:"PublicIp"`
	AllocationId *interface{} `json:"AllocationId"`
	State *interface{} `json:"State"`
	LineId *interface{} `json:"LineId"`
	BandWidth *interface{} `json:"BandWidth"`
	InstanceType *interface{} `json:"InstanceType"`
	InstanceId *interface{} `json:"InstanceId"`
	NetworkInterfaceId *interface{} `json:"NetworkInterfaceId"`
	BandWidthShareId *interface{} `json:"BandWidthShareId"`
	IsBandWidthShare *interface{} `json:"IsBandWidthShare"`
	IpVersion *interface{} `json:"IpVersion"`
	PassThrough *interface{} `json:"PassThrough"`
}

type DescribeAddressesOutput struct {
	_ struct{} `type:"structure"`
	RequestId *string `json:"RequestId"`
	AddressesSet []*Address `json:"AddressesSet"`
}

func (c *EIP) DescribeAddresses(input *DescribeAddressesInput) (map[string]interface{}, error) {
	req, _ := c.DescribeAddressesRequest(input)
	return req.Data.(map[string]interface{}),req.Send()
}

func (c *EIP) DescribeAddressesRequest(input *DescribeAddressesInput) (req *request.Request, output map[string]interface{}) {
	op := &request.Operation{
		Name:       "DescribeAddresses",
		HTTPMethod: "GET",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeAddressesInput{}
	}

	output = make(map[string]interface{})
	req = c.newRequest(op, input, output)
	return
}
