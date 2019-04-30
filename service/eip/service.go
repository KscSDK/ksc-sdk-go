package eip

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/ksc/ksc-sdk-go/ksc/kscquery"
)

type EIP struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "eip"   // Name of service.
	EndpointsID = ServiceName              // ID to lookup a service endpoint with.
	ServiceID   = "eip" // ServiceID is a unique identifer of a specific service.
)

func New(p client.ConfigProvider, cfgs ...*aws.Config) *EIP {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *EIP {
	svc := &EIP{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2016-03-04",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	//定制化builder
	svc.Handlers.Build.PushBackNamed(kscquery.BuildHandler)
	//如果需要返回json形式 使用这个handler
	svc.Handlers.Unmarshal.PushBackNamed(kscquery.UnmarshalHandler)
    svc.Handlers.UnmarshalMeta.PushBackNamed(kscquery.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(kscquery.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ELB operation and runs any
// custom request initialization.
func (c *EIP) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}

