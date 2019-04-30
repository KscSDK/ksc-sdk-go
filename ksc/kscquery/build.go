package kscquery

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/query/queryutil"
	"net/url"
	"strings"
)

// BuildHandler is a named request handler for building query protocol requests
var BuildHandler = request.NamedHandler{Name: "kscsdk.query.Build", Fn: Build}

// Build builds a request for an AWS Query service.
func Build(r *request.Request) {
	body := url.Values{
		"Action":  {r.Operation.Name},
		"Version": {r.ClientInfo.APIVersion},
	}
	if err := queryutil.Parse(body, r.Params, false); err != nil {
		r.Error = awserr.New("SerializationError", "failed encoding Query request", err)
		return
	}

	method := strings.ToUpper(r.HTTPRequest.Method)
	r.HTTPRequest.Header.Add("Accept", "application/json")

	if method == "GET" {
		r.HTTPRequest.URL.RawQuery = body.Encode()
	} else if method == "POST" {
		r.HTTPRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
		r.SetBufferBody([]byte(body.Encode()))
	} else if method == "PUT"{
		r.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
		r.SetBufferBody([]byte(body.Encode()))
	} else {
		r.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
		r.SetBufferBody([]byte(body.Encode()))
	}
}