package ksc

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"runtime"
)

// SDKVersionUserAgentHandler is a named request handler for building ksc agent requests
var SDKVersionUserAgentHandler = request.NamedHandler{
	Name: "Ksc.SDKVersionUserAgentHandler",
	Fn: request.MakeAddToUserAgentHandler(SDKName, SDKVersion,
		runtime.Version(), runtime.GOOS, runtime.GOARCH),
}
